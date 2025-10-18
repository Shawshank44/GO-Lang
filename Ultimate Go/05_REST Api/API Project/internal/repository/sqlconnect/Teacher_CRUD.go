package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/pkg/utils"
	"strconv"
)

func GetTeachersDBHandeler(teachers []models.Teacher, r *http.Request) ([]models.Teacher, error) {
	db, err := ConnectDB()

	if err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}

	defer rows.Close()

	// TeacherList := make([]models.Teacher, 0)
	for rows.Next() {
		var teacher models.Teacher
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error retrieving data")
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func GetTeacherDBHandeler(id int) (models.Teacher, error) {
	db, err := ConnectDB()

	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error retrieving data")
	}
	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {
		return models.Teacher{}, utils.ErrorHandler(err, "Error retrieving data")
	} else if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error retrieving data")
	}
	return teacher, nil
}

func POSTTeacherDBHandler(NewTeachers []models.Teacher) ([]models.Teacher, error) {
	db, err := ConnectDB()

	if err != nil {
		return nil, utils.ErrorHandler(err, "Error Posting data")
	}

	defer db.Close()

	stmt, err := db.Prepare(utils.GenerateInsertQuery("teachers", models.Teacher{}))
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error Posting data")
	}

	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(NewTeachers))
	for i, Newteacher := range NewTeachers {
		// res, err := stmt.Exec(Newteacher.FirstName, Newteacher.LastName, Newteacher.Email, Newteacher.Class, Newteacher.Subject)
		values := utils.GetStructValues(Newteacher)
		res, err := stmt.Exec(values...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error Posting data")
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error Posting data")
		}
		Newteacher.ID = int(lastID)
		addedTeachers[i] = Newteacher
	}
	return addedTeachers, nil
}

func PutTeacherDBHandler(id int, UpdatedTeacher models.Teacher) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {

		return models.Teacher{}, utils.ErrorHandler(err, "Error updating data")
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating data")
	} else if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating data")
	}
	UpdatedTeacher.ID = existingTeacher.ID
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", UpdatedTeacher.FirstName, UpdatedTeacher.LastName, UpdatedTeacher.Email, UpdatedTeacher.Class, UpdatedTeacher.Subject, UpdatedTeacher.ID)
	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating data")
	}
	return UpdatedTeacher, nil
}

func PatchTeachersDBHandler(updates []map[string]interface{}) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Error updating data")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Error updating data")
	}

	for _, update := range updates {
		idstr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			return utils.ErrorHandler(err, "Error updating data")
		}
		id, err := strconv.Atoi(idstr)
		if err != nil {
			tx.Rollback()

			return utils.ErrorHandler(err, "Error updating data")
		}

		var teacher models.Teacher
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(err, "Error updating data")
			}
			return utils.ErrorHandler(err, "Error updating data")
		}
		// Apply updates using reflect
		teacherVal := reflect.ValueOf(&teacher).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if k == "id" {
				continue // skip updating the id field
			}
			for i := 0; i < teacherVal.NumField(); i++ {
				field := teacherType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := teacherVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("Cannot convert %v to %v", val.Type(), fieldVal.Type())
							return utils.ErrorHandler(err, "Error updating data")
						}
					}
					break
				}
			}
		}
		_, err = tx.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", teacher.FirstName, teacher.LastName, teacher.Email, teacher.Class, teacher.Subject, teacher.ID)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "Error updating data")
		}
	}

	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "Error updating data")
	}
	return nil
}

func PatchTeacherDBHandler(id int, Updates map[string]interface{}) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating data")
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating data")
	} else if err != nil {
		log.Println(utils.ErrorHandler(err, "Error updating data"))
		return models.Teacher{

			// Apply Updates using Reflect package :
		}, utils.ErrorHandler(err, "Error updating data")
	}

	teacherVal := reflect.ValueOf(&existingTeacher).Elem()
	teacherType := teacherVal.Type()

	for k, v := range Updates {
		for i := 0; i < teacherVal.NumField(); i++ {
			field := teacherType.Field(i)
			if field.Tag.Get("json") == k+",omitempty" {
				if teacherVal.Field(i).CanSet() {
					teacherVal.Field(i).Set(reflect.ValueOf(v).Convert(teacherVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", existingTeacher.FirstName, existingTeacher.LastName, existingTeacher.Email, existingTeacher.Class, existingTeacher.Subject, existingTeacher.ID)
	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating data")
	}
	return existingTeacher, nil
}

func DeleteTeacherDBHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Error Deleting data")
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Error Deleting data")
	}
	fmt.Println(res.RowsAffected())
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Error Deleting data")
	}
	if rowsAffected == 0 {
		return utils.ErrorHandler(err, "Error Deleting data")
	}
	return nil
}

func DeleteTeachersDBHandler(ids []int) ([]int, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "Error Deleting data")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "Error Deleting data")
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = ?")
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}

	defer stmt.Close()

	deletedIds := []int{}

	for _, id := range ids {
		res, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return nil, utils.ErrorHandler(err, "Error Deleting data")
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, "Error Deleting data")
		}

		// if Teacher was deleted then add the ID to the deleted slice
		if rowsAffected > 0 {
			deletedIds = append(deletedIds, id)
		}
		if rowsAffected < 1 {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, "Error Deleting data")
		}
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "Error Deleting data")
	}

	if len(deletedIds) < 1 {
		return nil, utils.ErrorHandler(err, "Error Deleting data")
	}
	return deletedIds, nil
}

func GetStudentsByTeacherIDFromDB(teacherID string, students []models.Student) ([]models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}

	defer db.Close()

	query := `SELECT id, first_name, last_name, email, class FROM students WHERE class = (SELECT class from teachers WHERE id = ?)`
	rows, err := db.Query(query, teacherID)
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			log.Println(err)
			return nil, utils.ErrorHandler(err, "Error retrieving data")
		}
		students = append(students, student)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}
	return students, nil
}

func GetStudentsCountByTeacherIDFromDB(teacherID string) (int, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error retrieving data")
	}

	defer db.Close()

	query := `SELECT COUNT(*) FROM students WHERE class = (SELECT class FROM teachers WHERE id = ?)`

	var StudentCount int

	err = db.QueryRow(query, teacherID).Scan(&StudentCount)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error retrieving data")
	}
	return StudentCount, nil
}

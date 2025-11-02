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
	"strings"
)

func GetStudentsDBHandeler(students []models.Student, r *http.Request, limit, page int) ([]models.Student, int, error) {
	db, err := ConnectDB()

	if err != nil {
		return nil, 0, utils.ErrorHandler(err, "Error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class FROM students WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	// Add Pagination
	offset := (page - 1) * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, 0, utils.ErrorHandler(err, "Error retrieving data")
	}

	defer rows.Close()

	// StudentList := make([]models.Student, 0)
	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			return nil, 0, utils.ErrorHandler(err, "Error retrieving data")
		}
		students = append(students, student)
	}

	// Get the total count of students :
	var totalStudents int
	err = db.QueryRow("SELECT COUNT(*) FROM students").Scan(&totalStudents)
	if err != nil {
		utils.ErrorHandler(err, "")
		totalStudents = 0
	}

	return students, totalStudents, nil
}

func GetStudentDBHandeler(id int) (models.Student, error) {
	db, err := ConnectDB()

	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error retrieving data")
	}
	defer db.Close()

	var student models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "Error retrieving data")
	} else if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error retrieving data")
	}
	return student, nil
}

func POSTStudentDBHandler(NewStudents []models.Student) ([]models.Student, error) {
	db, err := ConnectDB()

	if err != nil {
		return nil, utils.ErrorHandler(err, "Error Posting data")
	}

	defer db.Close()

	stmt, err := db.Prepare(utils.GenerateInsertQuery("students", models.Student{}))
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error Posting data")
	}

	defer stmt.Close()

	addedStudents := make([]models.Student, len(NewStudents))
	for i, NewStudent := range NewStudents {
		// res, err := stmt.Exec(Newstudent.FirstName, Newstudent.LastName, Newstudent.Email, Newstudent.Class, Newstudent.Subject)
		values := utils.GetStructValues(NewStudent)
		res, err := stmt.Exec(values...)
		if err != nil {
			if strings.Contains(err.Error(), "a foreign key constraint fails (`school`.`students`, CONSTRAINT `students_ibfk_1` FOREIGN KEY (`class`) REFERENCES `teachers` (`class`))") {
				return nil, utils.ErrorHandler(err, "class/class teacher does not exists")
			}
			return nil, utils.ErrorHandler(err, "Error Posting data")
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error Posting data")
		}
		NewStudent.ID = int(lastID)
		addedStudents[i] = NewStudent
	}
	return addedStudents, nil
}

func PutStudentDBHandler(id int, UpdatedStudent models.Student) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {

		return models.Student{}, utils.ErrorHandler(err, "Error updating data")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&existingStudent.ID, &existingStudent.FirstName, &existingStudent.LastName, &existingStudent.Email, &existingStudent.Class)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "Error updating data")
	} else if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error updating data")
	}
	UpdatedStudent.ID = existingStudent.ID
	_, err = db.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", UpdatedStudent.FirstName, UpdatedStudent.LastName, UpdatedStudent.Email, UpdatedStudent.Class, UpdatedStudent.ID)
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error updating data")
	}
	return UpdatedStudent, nil
}

func PatchStudentsDBHandler(updates []map[string]interface{}) error {
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

		var student models.Student
		err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)

		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(err, "Error updating data")
			}
			return utils.ErrorHandler(err, "Error updating data")
		}
		// Apply updates using reflect
		studentVal := reflect.ValueOf(&student).Elem()
		studentType := studentVal.Type()

		for k, v := range update {
			if k == "id" {
				continue // skip updating the id field
			}
			for i := 0; i < studentVal.NumField(); i++ {
				field := studentType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := studentVal.Field(i)
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
		_, err = tx.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", student.FirstName, student.LastName, student.Email, student.Class, student.ID)
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

func PatchStudentDBHandler(id int, Updates map[string]interface{}) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return models.Student{}, utils.ErrorHandler(err, "Error updating data")
	}
	defer db.Close()

	var existingstudent models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&existingstudent.ID, &existingstudent.FirstName, &existingstudent.LastName, &existingstudent.Email, &existingstudent.Class)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "Error updating data")
	} else if err != nil {
		log.Println(utils.ErrorHandler(err, "Error updating data"))
		return models.Student{

			// Apply Updates using Reflect package :
		}, utils.ErrorHandler(err, "Error updating data")
	}

	studentVal := reflect.ValueOf(&existingstudent).Elem()
	studentType := studentVal.Type()

	for k, v := range Updates {
		for i := 0; i < studentVal.NumField(); i++ {
			field := studentType.Field(i)
			if field.Tag.Get("json") == k+",omitempty" {
				if studentVal.Field(i).CanSet() {
					studentVal.Field(i).Set(reflect.ValueOf(v).Convert(studentVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", existingstudent.FirstName, existingstudent.LastName, existingstudent.Email, existingstudent.Class, existingstudent.ID)
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error updating data")
	}
	return existingstudent, nil
}

func DeleteStudentDBHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Error Deleting data")
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM students WHERE id = ?", id)
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

func DeleteStudentsDBHandler(ids []int) ([]int, error) {
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

	stmt, err := tx.Prepare("DELETE FROM students WHERE id = ?")
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

		// if student was deleted then add the ID to the deleted slice
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

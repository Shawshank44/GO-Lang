package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"strconv"
	"strings"
)

func IsValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func IsValidSortField(field string) bool {
	validFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"class":      true,
		"subject":    true,
	}
	return validFields[field]
}

func AddSorting(r *http.Request, query string) string {
	sortParams := r.URL.Query()["sortby"] // converts in string slice

	if len(sortParams) > 0 {
		query += " ORDER BY"
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				continue
			}
			fields, order := parts[0], parts[1]
			if !IsValidSortField(fields) || !IsValidSortOrder(order) {
				continue
			}
			if i > 0 {
				query += ","
			}
			query += " " + fields + " " + order
		}
	}
	return query
}

func AddFilters(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for params, dbfield := range params {
		value := r.URL.Query().Get(params)
		if value != "" {
			query += " AND " + dbfield + " = ?"
			args = append(args, value)
		}
	}
	return query, args
}

func GetTeachersDBHandeler(teachers []models.Teacher, r *http.Request) ([]models.Teacher, error) {
	db, err := ConnectDB()

	if err != nil {
		// http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
	var args []interface{}

	query, args = AddFilters(r, query, args)

	query = AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Database Query error", http.StatusInternalServerError)
		return nil, err
	}

	defer rows.Close()

	// TeacherList := make([]models.Teacher, 0)
	for rows.Next() {
		var teacher models.Teacher
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			// http.Error(w, "Error scanning database results", http.StatusInternalServerError)
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func GetTeacherDBHandeler(id int) (models.Teacher, error) {
	db, err := ConnectDB()

	if err != nil {
		// http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {
		// http.Error(w, "Teacher not found", http.StatusNotFound)
		return models.Teacher{}, err
	} else if err != nil {
		// http.Error(w, "database query error", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return teacher, nil
}

func POSTTeacherDBHandler(NewTeachers []models.Teacher) ([]models.Teacher, error) {
	db, err := ConnectDB()

	if err != nil {
		// http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return nil, err
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?,?,?,?,?)")
	if err != nil {
		// http.Error(w, "Error in Preparing SQL query", http.StatusInternalServerError)
	}

	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(NewTeachers))
	for i, Newteacher := range NewTeachers {
		res, err := stmt.Exec(Newteacher.FirstName, Newteacher.LastName, Newteacher.Email, Newteacher.Class, Newteacher.Subject)
		if err != nil {
			// http.Error(w, "error inserting data in database", http.StatusInternalServerError)
			return nil, err
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			// http.Error(w, "Error getting last insert ID", http.StatusInternalServerError)
		}
		Newteacher.ID = int(lastID)
		addedTeachers[i] = Newteacher
	}
	return addedTeachers, nil
}

func PutTeacherDBHandler(id int, UpdatedTeacher models.Teacher) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		// log.Println(err)
		// http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		// http.Error(w, "User Not found or does not exists", http.StatusNotFound)
		return models.Teacher{}, err
	} else if err != nil {
		// log.Println(err)
		// http.Error(w, "Somthing went wrong unable to retreive", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	UpdatedTeacher.ID = existingTeacher.ID
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", UpdatedTeacher.FirstName, UpdatedTeacher.LastName, UpdatedTeacher.Email, UpdatedTeacher.Class, UpdatedTeacher.Subject, UpdatedTeacher.ID)
	if err != nil {
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return UpdatedTeacher, nil
}

func PatchTeachersDBHandler(updates []map[string]interface{}) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error Starting a transaction", http.StatusInternalServerError)
		return err
	}

	for _, update := range updates {
		idstr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			// http.Error(w, "Invalid teacher ID in updates", http.StatusBadRequest)
			return err
		}
		id, err := strconv.Atoi(idstr)
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Invalid teacher ID in updates", http.StatusBadRequest)
			return err
		}

		var teacher models.Teacher
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				// http.Error(w, "Teacher not found", http.StatusNotFound)
			}
			// http.Error(w, "error retrieving teacher", http.StatusInternalServerError)
			return err
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
							return err
						}
					}
					break
				}
			}
		}
		_, err = tx.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", teacher.FirstName, teacher.LastName, teacher.Email, teacher.Class, teacher.Subject, teacher.ID)
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		}
	}

	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error committ transaction", http.StatusInternalServerError)
	}
	return nil
}

func PatchTeacherDBHandler(id int, Updates map[string]interface{}) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		// http.Error(w, "User Not found or does not exists", http.StatusNotFound)
		return models.Teacher{}, err
	} else if err != nil {
		log.Println(err)
		// http.Error(w, "Somthing went wrong unable to retreive", http.StatusInternalServerError)
		return models.Teacher{

			// Apply Updates using Reflect package :
		}, err
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
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return existingTeacher, nil
}

func DeleteTeacherDBHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return err
	}
	fmt.Println(res.RowsAffected())
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error retrieving delete result", http.StatusInternalServerError)
		return err
	}
	if rowsAffected == 0 {
		// http.Error(w, "teacher not found", http.StatusNotFound)
		return err
	}
	return nil
}

func DeleteTeachersDBHandler(ids []int) ([]int, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error in connecting to Database", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error Starting a transaction", http.StatusInternalServerError)
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = ?")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		// http.Error(w, "Error preparing delete statement", http.StatusBadRequest)
	}

	defer stmt.Close()

	deletedIds := []int{}

	for _, id := range ids {
		res, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			// http.Error(w, "error deleting a teacher", http.StatusInternalServerError)
			return nil, err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error retrieving deleted result", http.StatusInternalServerError)
			return nil, err
		}

		// if Teacher was deleted then add the ID to the deleted slice
		if rowsAffected > 0 {
			deletedIds = append(deletedIds, id)
		}
		if rowsAffected < 1 {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("Id %d does not exists", id), http.StatusInternalServerError)
		}
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return nil, err
	}

	if len(deletedIds) < 1 {
		// http.Error(w, "IDs do not effect", http.StatusBadRequest)
		return nil, err
	}
	return deletedIds, nil
}

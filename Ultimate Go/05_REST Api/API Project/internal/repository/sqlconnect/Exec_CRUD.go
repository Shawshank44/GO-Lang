package sqlconnect

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/pkg/utils"
	"strconv"

	"golang.org/x/crypto/argon2"
)

func GetExecsDBHandeler(execs []models.Exec, r *http.Request) ([]models.Exec, error) {
	db, err := ConnectDB()

	if err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, username, user_created_at, inactive_status, role FROM Execs WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}

	defer rows.Close()

	// ExecList := make([]models.Exec, 0)
	for rows.Next() {
		var exec models.Exec
		err = rows.Scan(&exec.ID, &exec.FirstName, &exec.LastName, &exec.Email, &exec.Username, &exec.UserCreatedAt, &exec.InactiveStatus, &exec.Role)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error retrieving data")
		}
		execs = append(execs, exec)
	}
	return execs, nil
}

func GetExecDBHandeler(id int) (models.Exec, error) {
	db, err := ConnectDB()

	if err != nil {
		return models.Exec{}, utils.ErrorHandler(err, "Error retrieving data")
	}
	defer db.Close()

	var exec models.Exec
	err = db.QueryRow("SELECT id, first_name, last_name, email, username, inactive_status, role FROM execs WHERE id = ?", id).Scan(&exec.ID, &exec.FirstName, &exec.LastName, &exec.Email, &exec.Username, &exec.InactiveStatus, &exec.Role)
	if err == sql.ErrNoRows {
		return models.Exec{}, utils.ErrorHandler(err, "Error retrieving data")
	} else if err != nil {
		return models.Exec{}, utils.ErrorHandler(err, "Error retrieving data")
	}
	return exec, nil
}

func POSTExecDBHandler(NewExecs []models.Exec) ([]models.Exec, error) {
	db, err := ConnectDB()

	if err != nil {
		return nil, utils.ErrorHandler(err, "Error Posting data")
	}

	defer db.Close()

	stmt, err := db.Prepare(utils.GenerateInsertQuery("execs", models.Exec{}))
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error Posting data")
	}

	defer stmt.Close()

	addedExecs := make([]models.Exec, len(NewExecs))
	for i, NewExec := range NewExecs {
		if NewExec.Password == "" {
			return nil, utils.ErrorHandler(errors.New("please is blank"), "please enter password")
		}
		salt := make([]byte, 16)
		_, err := rand.Read(salt)
		if err != nil {
			return nil, utils.ErrorHandler(errors.New("failed to generate salt"), "error adding data")
		}
		hash := argon2.IDKey([]byte(NewExec.Password), salt, 1, 64*1024, 4, 32)
		saltBase64 := base64.StdEncoding.EncodeToString(salt)
		hashBase64 := base64.StdEncoding.EncodeToString(hash)

		encodedHash := fmt.Sprintf("%s.%s", saltBase64, hashBase64)

		NewExec.Password = encodedHash

		values := utils.GetStructValues(NewExec)
		res, err := stmt.Exec(values...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error Posting data")
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error Posting data")
		}
		NewExec.ID = int(lastID)
		addedExecs[i] = NewExec
	}
	return addedExecs, nil
}

func PatchExecsDBHandler(updates []map[string]interface{}) error {
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

		var execs models.Exec
		err = db.QueryRow("SELECT id, first_name, last_name, email, username FROM execs WHERE id = ?", id).Scan(&execs.ID, &execs.FirstName, &execs.LastName, &execs.Email, &execs.Username)

		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(err, "Error updating data")
			}
			return utils.ErrorHandler(err, "Error updating data")
		}
		// Apply updates using reflect
		execVal := reflect.ValueOf(&execs).Elem()
		execType := execVal.Type()

		for k, v := range update {
			if k == "id" {
				continue // skip updating the id field
			}
			for i := 0; i < execVal.NumField(); i++ {
				field := execType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := execVal.Field(i)
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
		_, err = tx.Exec("UPDATE execs SET first_name = ?, last_name = ?, email = ?, username = ? WHERE id = ?", execs.FirstName, execs.LastName, execs.Email, execs.Username, execs.ID)
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

func PatchExecDBHandler(id int, Updates map[string]interface{}) (models.Exec, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return models.Exec{}, utils.ErrorHandler(err, "Error updating data")
	}
	defer db.Close()

	var existingExecs models.Exec
	err = db.QueryRow("SELECT id, first_name, last_name, email, username FROM execs WHERE id = ?", id).Scan(&existingExecs.ID, &existingExecs.FirstName, &existingExecs.LastName, &existingExecs.Email, &existingExecs.Username)
	if err == sql.ErrNoRows {
		return models.Exec{}, utils.ErrorHandler(err, "Error updating data")
	} else if err != nil {
		log.Println(utils.ErrorHandler(err, "Error updating data"))
		return models.Exec{
			// Apply Updates using Reflect package :
		}, utils.ErrorHandler(err, "Error updating data")
	}

	ExecVal := reflect.ValueOf(&existingExecs).Elem()
	ExecType := ExecVal.Type()

	for k, v := range Updates {
		for i := 0; i < ExecVal.NumField(); i++ {
			field := ExecType.Field(i)
			if field.Tag.Get("json") == k+",omitempty" {
				if ExecVal.Field(i).CanSet() {
					ExecVal.Field(i).Set(reflect.ValueOf(v).Convert(ExecVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE execs SET first_name = ?, last_name = ?, email = ?, username = ? WHERE id = ?", existingExecs.FirstName, existingExecs.LastName, existingExecs.Email, existingExecs.Username, existingExecs.ID)
	if err != nil {
		return models.Exec{}, utils.ErrorHandler(err, "Error updating data")
	}
	return existingExecs, nil
}

func DeleteExecDBHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Error Deleting data")
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM execs WHERE id = ?", id)
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

func GetUserByUserName(username string) (*models.Exec, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer db.Close()

	user := &models.Exec{}
	err = db.QueryRow("SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs WHERE username = ?", username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.InactiveStatus, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorHandler(err, "user not found")
			return nil, utils.ErrorHandler(err, "user not found")
		}
		return nil, utils.ErrorHandler(err, "error in connecting database")
	}
	return user, nil
}

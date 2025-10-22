package handlers

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Get Execs
func GetExecutivesHandler(w http.ResponseWriter, r *http.Request) {
	var execs []models.Exec
	execs, err := sqlconnect.GetExecsDBHandeler(execs, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "Success",
		Count:  len(execs),
		Data:   execs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// Get Exec By ID
func GetExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exec, err := sqlconnect.GetExecDBHandeler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exec)
}

// POST Execs
func AddExecutivesHandler(w http.ResponseWriter, r *http.Request) {
	var NewExecs []models.Exec
	var RawExecs []map[string]interface{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &RawExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	fields := utils.GetFieldNames(models.Exec{})

	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	for _, exec := range RawExecs {
		for key := range exec {
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable fields found in request. Only use allowed fields", http.StatusBadRequest)
				return
			}
		}
	}

	err = json.Unmarshal(body, &NewExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	for _, exec := range NewExecs {
		err = utils.CheckBlankFields(exec)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedExecs, err := sqlconnect.POSTExecDBHandler(NewExecs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(addedExecs),
		Data:   addedExecs,
	}
	json.NewEncoder(w).Encode(response)
}

// PATCH Method (expects replacement of required data)
func PatchExecutivesHandler(w http.ResponseWriter, r *http.Request) {
	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchExecsDBHandler(updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PATCH Method by ID (expects replacement of required data)
func PatchExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Executive request", http.StatusBadRequest)
		return
	}
	var Updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&Updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}
	UpdatedExec, err := sqlconnect.PatchExecDBHandler(id, Updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UpdatedExec)

}

// DELETE Method
func DeleteExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Executive request", http.StatusBadRequest)
		return
	}
	err = sqlconnect.DeleteExecDBHandler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// w.WriteHeader(http.StatusNoContent)
	// Response Body
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Executive Sucessfully deleted",
		ID:     id,
	}

	json.NewEncoder(w).Encode(response)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Exec
	// Data Validation :
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}
	// Search for user if user actually exists :
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		utils.ErrorHandler(err, "error connecting to database")
		http.Error(w, "Error is connecting", http.StatusBadRequest)
		return
	}
	defer db.Close()

	user := &models.Exec{}
	err = db.QueryRow("SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs WHERE username = ?", req.Username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.InactiveStatus, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorHandler(err, "user not found")
			http.Error(w, "user not found.", http.StatusBadGateway)
			return
		}
		http.Error(w, "database connectivity error.", http.StatusBadGateway)
		return
	}

	// Is user active :
	if user.InactiveStatus {
		http.Error(w, "Account is inactive", http.StatusForbidden)
		return
	}

	// verify password :
	parts := strings.Split(user.Password, ".")
	if len(parts) != 2 {
		utils.ErrorHandler(errors.New("invalid encoded hash format"), "invalid encoded hash format")
		http.Error(w, "invalid encoded hash format", http.StatusForbidden)
		return
	}
	saltBase64 := parts[0]
	hashedPasswordBase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		utils.ErrorHandler(err, "failed to decode the salt")
		http.Error(w, "failed to decode the salt", http.StatusForbidden)
		return
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(hashedPasswordBase64)
	if err != nil {
		utils.ErrorHandler(err, "failed to decode the hash")
		http.Error(w, "failed to decode the hash", http.StatusForbidden)
		return
	}

	hash := argon2.IDKey([]byte(req.Password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		utils.ErrorHandler(errors.New("incorrect password"), "incorrect password")
		http.Error(w, "Incorrect password", http.StatusForbidden)
		return
	}

	if subtle.ConstantTimeCompare(hash, hashedPassword) == 1 {
		//do nothing
	} else {
		utils.ErrorHandler(errors.New("incorrect password"), "incorrect password")
		http.Error(w, "Incorrect password", http.StatusForbidden)
		return
	}

	// Generate Token :

	// Send token as a response or as a cookie
}

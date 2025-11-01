package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"
	"strconv"
	"time"
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
	user, err := sqlconnect.GetUserByUserName(req.Username)
	if err != nil {
		http.Error(w, "Error fetching User", http.StatusBadRequest)
		return
	}

	// Is user active :
	if user.InactiveStatus {
		http.Error(w, "Account is inactive", http.StatusForbidden)
		return
	}

	// verify password :
	err = utils.VerifyPassword(req.Password, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate JWT Token :
	tokenstring, err := utils.SignToken(user.ID, req.Username, user.Role)
	if err != nil {
		http.Error(w, "Could not create login token", http.StatusInternalServerError)
		return
	}
	// Send token as a response or as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    tokenstring,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "test",
	// 	Value:    "testvalue",
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	Expires:  time.Now().Add(24 * time.Hour),
	// 	SameSite: http.SameSiteStrictMode,
	// })

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenstring,
	}

	json.NewEncoder(w).Encode(response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message" : "User Logout Successsully"}`))
}

func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	userid, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid exec ID ", http.StatusBadRequest)
		return
	}

	var req models.UpdatePasswordRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	if req.CurrentPassword == "" || req.NewPassword == "" {
		http.Error(w, "fields cannot be blank", http.StatusBadRequest)
		return
	}

	token, err := sqlconnect.UpdatePasswordInDB(userid, req.CurrentPassword, req.NewPassword)
	if err != nil {
		http.Error(w, "Error in updating the password", http.StatusBadRequest)
		return
	}

	// Send token as a response or as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Message string
		Token   string `json:"token"`
	}{
		Message: "Password updated successfully",
		Token:   token,
	}

	json.NewEncoder(w).Encode(response)
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	if req.Email == "" {
		http.Error(w, "email field cannot be blank!", http.StatusBadRequest)
	}

	err = sqlconnect.ForgotPasswordDBHandler(req.Email)
	if err != nil {
		return
	}

	// respond with a Success confirmation :
	fmt.Fprintf(w, "Password sent link sent to %s", req.Email)

}

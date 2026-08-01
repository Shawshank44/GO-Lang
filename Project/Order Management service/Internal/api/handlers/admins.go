package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order_mgt/Internal/api/middlewares"
	"order_mgt/Internal/models"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	"order_mgt/pkg/utils"
	utilssql "order_mgt/pkg/utils_sql"
	"strings"
	"time"
)

func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allwoed", http.StatusBadRequest)
		return
	}

	var req models.Admin

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
		http.Error(w, "Fields cannot be empty", http.StatusBadRequest)
		return
	}

	exists, err := utilssql.EmailExists(r.Context(), req.Email)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "email already exists", http.StatusConflict)
		return
	}

	userID, err := sqlconnect.RegisterAdminToDB(r.Context(), req)
	if err != nil {
		http.Error(w, "Unable to create your account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := struct {
		Status string `json:"status"`
		ID     int64  `json:"id"`
	}{
		Status: "User Successfully Created",
		ID:     userID,
	}

	json.NewEncoder(w).Encode(res)
}

func GetAdmins(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	page, limit := utils.GetPaginationParams(r)

	userList, totalUsers, err := sqlconnect.GetAdminsFromDB(r.Context(), r, limit, page)
	if err != nil {
		http.Error(w, "Unable to fetch admins", http.StatusBadRequest)
		return
	}

	totalPages := (totalUsers + limit - 1) / limit

	res := struct {
		Status     string
		Count      int
		TotalPages int
		PageNo     int
		PageSize   int
		Data       []models.AdminResponse
	}{
		Status:     "Success",
		Count:      totalUsers,
		TotalPages: totalPages,
		PageNo:     page,
		PageSize:   limit,
		Data:       userList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uid, ok := r.Context().Value(middlewares.UserIDkey).(float64)
	if !ok {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}

	id := int(uid)

	admin, err := sqlconnect.GetAdminFromDB(r.Context(), id)
	if err != nil {
		http.Error(w, "Unable to fetch the user details", http.StatusBadRequest)
		return
	}

	res := struct {
		Status string
		Data   models.AdminResponse
	}{
		Status: "Success",
		Data:   admin,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.Admin

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		http.Error(w, "Fields cannot be empty", http.StatusBadRequest)
		return
	}

	admin, err := sqlconnect.LoginAdminFromDB(r.Context(), req.Username)
	if err != nil {
		http.Error(w, "Error fetching the user", http.StatusBadRequest)
		return
	}

	if admin.InactiveStatus {
		http.Error(w, "User inactive kindly contact adminstrator", http.StatusForbidden)
		return
	}

	err = utils.VerifyPassword(req.Password, admin.Password)
	if err != nil {
		http.Error(w, "user does not exists", http.StatusInternalServerError)
		return
	}

	tokenString, err := utils.SignToken(admin.ID, admin.Username, "admin")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	res := struct {
		UserID int
		Token  string `json:"token"`
	}{
		UserID: admin.ID,
		Token:  tokenString,
	}

	json.NewEncoder(w).Encode(res)
}

func LogoutAdmin(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte(`{"Message" : "User Logout Successfully"}`))
}

func UpdateAdminDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uid, ok := r.Context().Value(middlewares.UserIDkey).(float64)
	if !ok {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}
	id := int(uid)

	var req models.AdminUpdateDetail
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Inavlid json payload", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Email) == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	exists, err := utilssql.EmailExists(r.Context(), req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "email already exists", http.StatusConflict)
		return
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = sqlconnect.UpdateAdminDetailsInDB(r.Context(), otp, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.SendOTPEmail(req.Email, otp, "Your Email change request - Orderfy.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := struct {
		Success bool
		Message string
	}{
		Success: true,
		Message: fmt.Sprintf("Email change request has been shared to %s", req.Email),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func ConfirmAdminDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	uid, ok := r.Context().Value(middlewares.UserIDkey).(float64)
	if !ok {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}
	id := int(uid)

	var req models.ConfirmDetailAdmins
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Otp) == "" || strings.TrimSpace(req.Email) == "" {
		http.Error(w, "otp is required", http.StatusBadRequest)
		return
	}

	exists, err := utilssql.EmailExists(r.Context(), req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "Email already exists", http.StatusInternalServerError)
		return
	}

	err = sqlconnect.ConfirmAdminDetailsInDB(r.Context(), req, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := struct {
		Success bool
		Message string
	}{
		Success: true,
		Message: fmt.Sprintf("email address has been updated to %s", req.Email),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

func DeactivateAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uid, ok := r.Context().Value(middlewares.UserIDkey).(float64)
	if !ok {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}
	id := int(uid)

	err := sqlconnect.DeactivateAdminFromDB(r.Context(), id)
	if err != nil {
		http.Error(w, "Unable to deactivate user from DB", http.StatusInternalServerError)
		return
	}

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
	res := struct {
		Status string
		ID     int
	}{
		Status: "User Successfully deactivated",
		ID:     id,
	}

	json.NewEncoder(w).Encode(res)

}

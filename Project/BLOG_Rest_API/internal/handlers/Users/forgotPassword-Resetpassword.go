package users

import (
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Users_SQL"
	utilssql "blog_rest_api/internal/repositories/Utils_SQL"
	"blog_rest_api/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not allowed", http.StatusBadRequest)
		return
	}

	var req models.UserUpdateDetail
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Email) == "" {
		http.Error(w, "Email field cannot be blank", http.StatusBadRequest)
		return
	}

	exists, err := utilssql.EmailExists(r.Context(), req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Invalid email", http.StatusConflict)
		return
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.ForgorPasswordFromDB(r.Context(), otp, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.SendOTPEmail(req.Email, otp, "Your Password change request - Blogsup.com")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := struct {
		Success bool
		Message string
	}{
		Success: true,
		Message: fmt.Sprintf("Password change request has been shared to %s", req.Email),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not allowed", http.StatusBadRequest)
		return
	}

	var req models.UpdatePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Otp) == "" || strings.TrimSpace(req.NewPassword) == "" {
		http.Error(w, "Email field cannot be blank", http.StatusBadRequest)
		return
	}

	err = repositories.ResetPasswordFromDB(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := struct {
		Success bool
		Message string
	}{
		Success: true,
		Message: fmt.Sprintln("Password has been updated kindly login."),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

package users

import (
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Users_SQL"
	utilssql "blog_rest_api/internal/repositories/Utils_SQL"
	"blog_rest_api/internal/services"
	"blog_rest_api/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func UpdateDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authID, err := services.UserAuthService(r.Context(), r)
	if err != nil {
		http.Error(w, "Unable to get the user Id from JWT", http.StatusUnauthorized)
		return
	}

	if idStr != authID {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	var req models.UserUpdateDetail
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, "email already exists.", http.StatusConflict)
		return
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.UpdateDetailsInDB(r.Context(), otp, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.SendOTPEmail(req.Email, otp, "Your Email change request - Blogsup.com")
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

func Confirmdetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authID, err := services.UserAuthService(r.Context(), r)
	if err != nil {
		http.Error(w, "Unable to get the user Id from JWT", http.StatusUnauthorized)
		return
	}

	if idStr != authID {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	var req models.ConfirmDetail
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Otp) == "" {
		http.Error(w, "otp is required", http.StatusBadRequest)
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
		http.Error(w, "email already exists.", http.StatusConflict)
		return
	}

	err = repositories.ConfirmDetailsInDB(r.Context(), req, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := struct {
		Success bool
		Message string
	}{
		Success: true,
		Message: fmt.Sprintf("email addess has been updated to %s", req.Email),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

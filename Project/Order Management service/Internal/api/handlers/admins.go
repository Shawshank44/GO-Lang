package handlers

import (
	"encoding/json"
	"net/http"
	"order_mgt/Internal/models"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	"order_mgt/pkg/utils"
	utilssql "order_mgt/pkg/utils_sql"
	"strconv"
	"strings"
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

	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

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

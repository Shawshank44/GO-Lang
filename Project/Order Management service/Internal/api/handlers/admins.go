package handlers

import (
	"encoding/json"
	"net/http"
	"order_mgt/Internal/models"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	utilssql "order_mgt/pkg/utils_sql"
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

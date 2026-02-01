package users

import (
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Users_SQL"
	"encoding/json"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Fields cannot be blank", http.StatusBadRequest)
	}

	exists, err := repositories.EmailExists(r.Context(), req)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "email already exists.", http.StatusConflict)
	}

	userID, err := repositories.RegisterUserToDB(r.Context(), req)
	if err != nil {
		http.Error(w, "Unable to create your account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := struct {
		Status string `json:"status"`
		ID     int64
	}{
		Status: "User Successfully created",
		ID:     userID,
	}

	json.NewEncoder(w).Encode(res)

}

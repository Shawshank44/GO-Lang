package users

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func DeactivateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Unable to convert Teacher id ", http.StatusBadRequest)
		return
	}

	authid, err := services.UserAuthService(r.Context(), r)
	if err != nil {
		http.Error(w, "Inavlid user id", http.StatusBadRequest)
		return
	}

	if authid != idstr {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	res, err := db.ExecContext(r.Context(), "DELETE FROM users WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Error in deleting user from database", http.StatusInternalServerError)
		return
	}
	fmt.Println(res.RowsAffected())

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "rows affected 0", http.StatusInternalServerError)
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
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "User Successfully deactivated",
		ID:     id,
	}

	json.NewEncoder(w).Encode(response)
}

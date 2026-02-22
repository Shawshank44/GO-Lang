package users

import (
	repositories "blog_rest_api/internal/repositories/Users_SQL"
	"blog_rest_api/internal/services"
	"encoding/json"
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

	err = repositories.DeactivateUserFromDB(r.Context(), id)
	if err != nil {
		http.Error(w, "Unable to Deactivate user fromDB", http.StatusInternalServerError)
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

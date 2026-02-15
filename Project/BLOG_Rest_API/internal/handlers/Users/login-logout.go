package users

import (
	"blog_rest_api/internal/auth"
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Users_SQL"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		http.Error(w, "Fields cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := repositories.LoginDB(r.Context(), req.Username)
	if err != nil {
		http.Error(w, "Error fetching user", http.StatusBadRequest)
		return
	}

	if user.InactiveStatus {
		http.Error(w, "Account is inactive", http.StatusForbidden)
		return
	}

	err = auth.VerifyPassword(req.Password, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenString, err := auth.SignToken(user.ID, req.Username, "client")
	if err != nil {
		http.Error(w, "Could not create login session", http.StatusInternalServerError)
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
		Token string `json:"token"`
	}{
		Token: tokenString,
	}

	json.NewEncoder(w).Encode(res)
}

func Logout(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte(`{"message" : "User Logout Successfully"}`))
}

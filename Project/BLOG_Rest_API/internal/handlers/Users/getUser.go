package users

import (
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Users_SQL"
	"blog_rest_api/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userList, err := repositories.GetUsersFromDB(r.Context(), r)
	if err != nil {
		http.Error(w, "Unable to fetch users", http.StatusBadRequest)
		return
	}
	res := struct {
		Status string
		Count  int
		Data   []models.UserResponse
	}{
		Status: "Success",
		Count:  len(userList),
		Data:   userList,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	authID, err := services.UserAuthService(r.Context(), r)
	if err != nil {
		http.Error(w, "Unable to get the user Id from JWT", http.StatusUnauthorized)
		return
	}

	if idstr != authID {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	user, err := repositories.GetUserFromDB(r.Context(), id)
	if err != nil {
		http.Error(w, "Unable to fetch the user details", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

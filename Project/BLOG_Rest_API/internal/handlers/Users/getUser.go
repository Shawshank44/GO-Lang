package users

import (
	"blog_rest_api/internal/models"
	repositories "blog_rest_api/internal/repositories/Users_SQL"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userList, err := repositories.GetUsersFromDB(r)
	if err != nil {
		http.Error(w, "Unable to fetch users", http.StatusBadRequest)
	}
	res := struct {
		Status string
		Count  int
		Data   []models.UserResponse
	}{
		Status: "Sucess",
		Count:  len(userList),
		Data:   userList,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

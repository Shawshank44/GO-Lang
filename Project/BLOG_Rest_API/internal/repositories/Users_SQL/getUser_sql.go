package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/internal/models"
	"blog_rest_api/pkg/utils"
	"net/http"
)

func GetUsersFromDB(r *http.Request) ([]models.UserResponse, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to connect to server")
	}
	defer db.Close()

	query := "SELECT id, username, email, user_created_at, password_changed_at, inactive_status FROM users WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)
	query = utils.AddSorting(r, query)

	rows, err := db.QueryContext(r.Context(), query, args...)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to query the user")
	}
	defer rows.Close()

	userList := make([]models.UserResponse, 0)
	for rows.Next() {
		var user models.UserResponse
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.UserCreatedAt, &user.PasswordChangedAT, &user.InactiveStatus)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Unable to find the roq")
		}
		userList = append(userList, user)
	}
	return userList, nil
}

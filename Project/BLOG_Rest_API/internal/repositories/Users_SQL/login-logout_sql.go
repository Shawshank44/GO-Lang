package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/internal/models"
	"blog_rest_api/pkg/utils"
	"context"
	"database/sql"
)

func LoginDB(ctx context.Context, username string) (*models.User, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Server Error")
	}

	defer db.Close()

	user := &models.User{}
	err = db.QueryRowContext(ctx, "SELECT id, username, email, password, inactive_status FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.InactiveStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrorHandler(err, "user not found")
		}
		return nil, utils.ErrorHandler(err, "error in connecting to database")
	}
	return user, nil
}

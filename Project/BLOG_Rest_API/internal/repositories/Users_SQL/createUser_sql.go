package repositories

import (
	"blog_rest_api/internal/auth"
	"blog_rest_api/internal/db"
	"blog_rest_api/internal/models"
	"blog_rest_api/pkg/utils"
	"context"
)

func EmailExists(ctx context.Context, req models.User) (bool, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return false, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)`

	var exists bool
	err = db.QueryRowContext(ctx, query, req.Email).Scan(&exists)
	return exists, err

}

func RegisterUserToDB(ctx context.Context, req models.User) (int64, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	hashedpassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error in hashing the password")
	}

	result, err := db.ExecContext(ctx, query, req.Username, req.Email, hashedpassword)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error in quering the database")
	}

	userID, _ := result.LastInsertId()
	return userID, nil
}

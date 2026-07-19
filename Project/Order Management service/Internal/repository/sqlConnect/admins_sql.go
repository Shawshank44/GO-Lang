package sqlconnect

import (
	"context"
	"order_mgt/Internal/models"
	"order_mgt/pkg/utils"
)

func RegisterAdminToDB(ctx context.Context, req models.Admin) (int64, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	query := `INSERT INTO admins (username, email, password) VALUES (?, ?, ?)`

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error in hashing the password")
	}

	result, err := db.ExecContext(ctx, query, req.Username, req.Email, hashedPassword)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error in quering the database")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error in quering the last inserted ID")
	}

	return userID, nil

}

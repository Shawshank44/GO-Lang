package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/internal/models"
	"blog_rest_api/pkg/utils"
	"context"
	"time"
)

func UpdateDetailsInDB(ctx context.Context, otp string, id int) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	mins := time.Duration(10)
	expiry := time.Now().Add(mins * time.Minute).Format(time.RFC3339)

	_, err = db.ExecContext(ctx, "UPDATE users SET password_otp = ?, otp_expires = ? WHERE id = ?", otp, expiry, id)
	if err != nil {
		return utils.ErrorHandler(err, "unable to update the users in DB")
	}
	return nil
}

func ConfirmDetailsInDB(ctx context.Context, req models.ConfirmDetail, id int) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	var userID int

	query := "SELECT id FROM users WHERE password_otp = ? AND otp_expires > ?"

	err = db.QueryRowContext(ctx, query, req.Otp, time.Now().Format(time.RFC3339)).Scan(&userID)
	if err != nil {
		return utils.ErrorHandler(err, "Otp is invalid or either expired")
	}

	_, err = db.ExecContext(ctx, "UPDATE users SET email = ?, password_otp = NULL, otp_expires = NULL WHERE id = ?", req.Email, id)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to update the email id in DB")
	}
	return nil
}

package repositories

import (
	"blog_rest_api/internal/auth"
	"blog_rest_api/internal/db"
	"blog_rest_api/internal/models"
	"blog_rest_api/pkg/utils"
	"context"
	"errors"
	"time"
)

func ForgorPasswordFromDB(ctx context.Context, otp, email string) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	mins := time.Duration(10)
	expiry := time.Now().Add(mins * time.Minute).Format(time.RFC3339)

	_, err = db.ExecContext(ctx, "UPDATE users SET password_otp = ?, otp_expires = ? WHERE email = ?", otp, expiry, email)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to update the user fields in DB")
	}

	return nil
}

func ResetPasswordFromDB(ctx context.Context, req models.UpdatePasswordRequest) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	var userID int
	var currentPassword string

	query := "SELECT id, password FROM users WHERE password_otp = ? AND otp_expires > ?"

	err = db.QueryRowContext(ctx, query, req.Otp, time.Now().Format(time.RFC3339)).Scan(&userID, &currentPassword)
	if err != nil {
		return utils.ErrorHandler(err, "Invalid OTP or expired")
	}

	same, err := auth.IsSameAsOldPassword(req.NewPassword, currentPassword)
	if err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	if same {
		return utils.ErrorHandler(errors.New("new password must be different from old password"), "new password must be different from old password")
	}

	newpassword, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to hash the new password")
	}

	updateQuery := "UPDATE users SET password = ?, password_otp = NULL, otp_expires = NULL, password_changed_at = CURRENT_TIMESTAMP WHERE id = ?"
	_, err = db.ExecContext(ctx, updateQuery, newpassword, userID)
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	return nil
}

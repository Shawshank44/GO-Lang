package sqlconnect

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"order_mgt/Internal/models"
	"order_mgt/pkg/utils"
	"time"
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

func GetAdminsFromDB(ctx context.Context, r *http.Request, limit, page int) ([]models.AdminResponse, int, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, 0, utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	query := "SELECT id, username, email, user_created_at, password_changed_at, inactive_status FROM admins WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	query = utils.AddSorting(r, query)

	// Pagination :
	offset := (page - 1) * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, 0, utils.ErrorHandler(err, "Unable to query the admin")
	}

	defer rows.Close()

	adminList := make([]models.AdminResponse, 0)

	for rows.Next() {
		var admin models.AdminResponse
		err = rows.Scan(&admin.ID, &admin.Username, &admin.Email, &admin.UserCreatedAt, &admin.PasswordChangedAt, &admin.InactiveStatus)
		if err == sql.ErrNoRows {
			return nil, 0, nil
		}
		if err != nil {
			return nil, 0, utils.ErrorHandler(err, "Unable to find the row")
		}
		adminList = append(adminList, admin)
	}

	var totalusers int
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM admins").Scan(&totalusers)
	if err != nil {
		utils.ErrorHandler(err, "")
		totalusers = 0
	}

	return adminList, totalusers, nil
}

func GetAdminFromDB(ctx context.Context, id int) (models.AdminResponse, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.AdminResponse{}, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	var admin models.AdminResponse

	query := "SELECT id, username, email, user_created_at, password_changed_at, inactive_status FROM admins WHERE id = ?"

	err = db.QueryRowContext(ctx, query, id).Scan(&admin.ID, &admin.Username, &admin.Email, &admin.UserCreatedAt, &admin.PasswordChangedAt, &admin.InactiveStatus)
	if err == sql.ErrNoRows {
		return models.AdminResponse{}, utils.ErrorHandler(err, "unable to find the user")
	}
	if err != nil {
		return models.AdminResponse{}, utils.ErrorHandler(err, "unable to find the user")
	}

	return admin, nil

}

func LoginAdminFromDB(ctx context.Context, username string) (*models.Admin, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	user := &models.Admin{}

	err = db.QueryRowContext(ctx, "SELECT id, username, email, password, inactive_status FROM admins WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.InactiveStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrorHandler(err, "user not found")
		}
		return nil, utils.ErrorHandler(err, "error in connecting to database")
	}
	return user, nil
}

func UpdateAdminDetailsInDB(ctx context.Context, otp string, id int) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	mins := time.Duration(10)
	expiry := time.Now().Add(mins * time.Minute).Format(time.RFC3339)

	_, err = db.ExecContext(ctx, "UPDATE admins SET password_otp = ?, otp_expires = ? WHERE id = ?", otp, expiry, id)
	if err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	return nil
}

func ConfirmAdminDetailsInDB(ctx context.Context, req models.ConfirmDetailAdmins, id int) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	var userID int

	query := "SELECT id FROM admins WHERE password_otp = ? AND otp_expires > ?"
	err = db.QueryRowContext(ctx, query, req.Otp, time.Now().Format(time.RFC3339)).Scan(&userID)
	if err != nil {
		return utils.ErrorHandler(err, "Otp is invalid or either expired.")
	}

	_, err = db.ExecContext(ctx, "UPDATE admins SET email = ?, password_otp = NULL, otp_expires = NULL WHERE id = ?", req.Email, id)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to update the email id in DB")
	}

	return nil
}

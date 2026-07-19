package sqlconnect

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
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

package utilssql

import (
	"context"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	"order_mgt/pkg/utils"
)

func EmailExists(ctx context.Context, email string) (bool, error) {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		return false, utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	query := `SELECT EXISTS (SELECT 1 FROM admins WHERE email = ?)`

	var exists bool
	err = db.QueryRowContext(ctx, query, email).Scan(&exists)

	return exists, err
}

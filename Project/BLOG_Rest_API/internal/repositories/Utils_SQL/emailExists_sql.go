package utilssql

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/pkg/utils"
	"context"
)

func EmailExists(ctx context.Context, email string) (bool, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return false, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)`

	var exists bool
	err = db.QueryRowContext(ctx, query, email).Scan(&exists)
	return exists, err
}

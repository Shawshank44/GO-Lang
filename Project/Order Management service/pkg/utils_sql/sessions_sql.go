package utilssql

import (
	"context"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	"order_mgt/pkg/utils"
)

func CreateSessionsInDB(ctx context.Context, sessionID string) error {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	_, err = db.ExecContext(ctx, "INSERT INTO upload_sessions(id) VALUES(?)", sessionID)
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	return nil
}

func ValidateSessionsInDB(ctx context.Context, sessionID string) (bool, error) {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		return false, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	query := `SELECT EXISTS (SELECT 1 FROM upload_sessions WHERE id = ?)`

	var exists bool

	err = db.QueryRowContext(ctx, query, sessionID).Scan(&exists)

	return exists, err
}

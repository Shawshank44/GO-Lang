package utilssql

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/pkg/utils"
	"context"
)

func CreateSessionInDB(ctx context.Context, sessionID string) error {
	db, err := db.ConnectDB()
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

func FinalizeSessionUploads(ctx context.Context, sessionID string) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	_, err = db.ExecContext(ctx, "UPDATE uploads SET session_id=NULL WHERE session_id = ?", sessionID)
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error unable to query session.")
	}

	return nil
}

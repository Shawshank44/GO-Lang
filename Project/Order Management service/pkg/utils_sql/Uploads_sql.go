package utilssql

import (
	"context"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	"order_mgt/pkg/utils"
)

func UploadToDB(ctx context.Context, sessionID string, fileURL string) error {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	_, err = db.ExecContext(ctx, "INSERT INTO files(file_path,session_id) VALUES(?,?)", fileURL, sessionID)
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}
	return nil
}

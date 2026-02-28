package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/pkg/utils"
	"context"
)

func UploadToDB(ctx context.Context, userId int, fileURL, contentType string) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	_, err = db.ExecContext(ctx, "INSERT INTO assets (user_id, file_url, file_type) VALUES (?, ?, ?)", userId, fileURL, contentType)
	if err != nil {
		return utils.ErrorHandler(err, "Insertion failed Internal server error")
	}
	return nil
}

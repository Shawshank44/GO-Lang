package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/pkg/utils"
	"context"
	"encoding/json"
)

func CreatePostInDB(ctx context.Context, username string, title string, content json.RawMessage, tags []byte) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, "INSERT INTO posts (username, title, content, tags) VALUES (?,?,?,?)", username, title, content, tags)
	if err != nil {
		return utils.ErrorHandler(err, err.Error())
	}

	return nil
}

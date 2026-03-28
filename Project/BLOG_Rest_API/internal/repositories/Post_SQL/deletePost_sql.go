package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/pkg/utils"
	"context"
	"encoding/json"
	"os"
)

func DeletePostFromDB(ctx context.Context, id int) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	var contentJSON []byte

	err = db.QueryRowContext(ctx, "SELECT content FROM posts WHERE id=?", id).Scan(&contentJSON)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to find the content")
	}

	content := make(map[string]interface{}, 0)

	err = json.Unmarshal(contentJSON, &content)
	if err != nil {
		return utils.ErrorHandler(err, "Invalid Json payload")
	}

	images := utils.ExtractImages(content)

	for _, i := range images {
		err = os.Remove(`C:\Users\Shashank.BR\OneDrive\Desktop\Go programing\Project\BLOG_Rest_API\cmd\server\` + i)
		if err != nil {
			return utils.ErrorHandler(err, "Umable to delete the images from DB")
		}

		_, err = db.ExecContext(ctx, "DELETE FROM uploads WHERE file_path=?", i)
		if err != nil {
			return utils.ErrorHandler(err, "Internal server error unable to delete images")
		}
	}

	_, err = db.ExecContext(ctx, "DELETE FROM posts WHERE id=?", id)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to delete the post.")
	}

	return nil
}

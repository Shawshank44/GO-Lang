package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/internal/models"
	"blog_rest_api/pkg/utils"
	"context"
	"encoding/json"
	"os"
)

func deleteRemovedImages(ctx context.Context, oldImgs, newImgs []string) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	newSet := map[string]struct{}{}

	for _, img := range newImgs {
		newSet[img] = struct{}{}
	}

	for _, old := range oldImgs {
		if _, exists := newSet[old]; !exists {
			err = os.Remove(`C:\Users\Shashank.BR\OneDrive\Desktop\Go programing\Project\BLOG_Rest_API\cmd\server\` + old)
			if err != nil {
				return utils.ErrorHandler(err, "Internal error unable to delete the files in DB")
			}

			_, err := db.ExecContext(ctx, "DELETE FROM uploads WHERE file_path=?", old)
			if err != nil {
				return utils.ErrorHandler(err, "Internal error unable to delete the files in DB")
			}
		}
	}
	return nil
}

func UpdatePostInDB(ctx context.Context, post *models.Post, id int) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return utils.ErrorHandler(err, "Internal transact error")
	}

	defer tx.Rollback()

	var oldContentJSON []byte

	err = tx.QueryRowContext(ctx, `SELECT content FROM posts WHERE id=?`, id).Scan(&oldContentJSON)
	if err != nil {
		return utils.ErrorHandler(err, "Post not found")
	}

	var oldContent map[string]interface{}
	var newcontent map[string]interface{}

	err = json.Unmarshal(oldContentJSON, &oldContent)
	if err != nil {
		return utils.ErrorHandler(err, "Inavalid old content")
	}

	err = json.Unmarshal(oldContentJSON, &oldContent)
	if err != nil {
		return utils.ErrorHandler(err, "Inavalid old content")
	}

	err = json.Unmarshal(post.Content, &newcontent)
	if err != nil {
		return utils.ErrorHandler(err, "Invalid new content")
	}

	oldImages := utils.ExtractImages(oldContent)
	newImages := utils.ExtractImages(newcontent)

	contentJSON, err := json.Marshal(newcontent)
	if err != nil {
		return utils.ErrorHandler(err, "content serialization failed")
	}

	tags, err := json.Marshal(post.Tags)
	if err != nil {
		return utils.ErrorHandler(err, "invalid payload")
	}

	res, err := tx.ExecContext(ctx, `UPDATE posts SET title=?, content=?, tags=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`, post.Title, contentJSON, tags, id)
	if err != nil {
		return utils.ErrorHandler(err, "update failed")
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return utils.ErrorHandler(err, "post not found")
	}

	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "transaction failed")
	}

	err = deleteRemovedImages(ctx, oldImages, newImages)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to delete the old and New images")
	}

	return nil
}

package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/internal/models"
	"blog_rest_api/pkg/utils"
	"context"
	"database/sql"
	"encoding/json"
)

func GetPostsFromDB(ctx context.Context) (*[]models.Post, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	query := `SELECT id, username, title, content, tags, created_at, updated_at FROM posts ORDER BY created_at`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to fetch the posts")
	}

	defer rows.Close()

	posts := make([]models.Post, 0)

	for rows.Next() {
		var post models.Post
		var tags string

		err := rows.Scan(&post.ID, &post.Username, &post.Title, &post.Content, &tags, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Unable to scan the post")
		}

		if tags != "" {
			err = json.Unmarshal([]byte(tags), &post.Tags)
			if err != nil {
				return nil, err
			}
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, utils.ErrorHandler(err, "Error Iterating the posts")
	}

	return &posts, nil

}

func MyPostsFromDB(ctx context.Context, username string) (*[]models.Post, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal server error")
	}
	defer db.Close()

	query := `SELECT id, username, title, content, tags, created_at, updated_at FROM posts WHERE username = ?`

	rows, err := db.QueryContext(ctx, query, username)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to fetch the posts")
	}

	defer rows.Close()

	posts := make([]models.Post, 0)

	for rows.Next() {
		var post models.Post
		var tags string

		rows.Scan(&post.ID, &post.Username, &post.Title, &post.Content, &tags, &post.CreatedAt, &post.UpdatedAt)

		if err != nil {
			return nil, utils.ErrorHandler(err, "Unable to scan the post")
		}

		if tags != "" {
			err = json.Unmarshal([]byte(tags), &post.Tags)
			if err != nil {
				return nil, err
			}
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, utils.ErrorHandler(err, "Error Iterating the posts")
	}

	return &posts, nil
}

func GetPostFromDB(ctx context.Context, id int) (*models.Post, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	var post models.Post
	var tags string

	query := `SELECT id, username, title, content, tags, created_at, updated_at FROM posts WHERE id = ?`

	err = db.QueryRowContext(ctx, query, id).Scan(&post.ID, &post.Username, &post.Title, &post.Content, &tags, &post.CreatedAt, &post.UpdatedAt)
	if err == sql.ErrNoRows {
		return &models.Post{}, utils.ErrorHandler(err, "Unable to find the post from db")
	}
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to find the post from db")
	}

	if tags != "" {
		err = json.Unmarshal([]byte(tags), &post.Tags)
		if err != nil {
			return nil, err
		}
	}

	return &post, nil
}

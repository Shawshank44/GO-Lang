package repositories

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/pkg/utils"
	"context"
	"fmt"
)

func DeactivateUserFromDB(ctx context.Context, id int) error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal serve error")
	}

	defer db.Close()

	res, err := db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return utils.ErrorHandler(err, "Internal serve error")
	}
	fmt.Println(res.RowsAffected())

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "Internal serve error")
	}

	if rowsAffected == 0 {
		return utils.ErrorHandler(err, "Internal serve error")
	}
	return nil
}

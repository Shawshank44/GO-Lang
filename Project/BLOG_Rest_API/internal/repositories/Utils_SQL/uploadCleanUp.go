package utilssql

import (
	"blog_rest_api/internal/db"
	"blog_rest_api/pkg/utils"
	"log"
	"os"
)

func UploadCleanUp() error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	rows, err := db.Query(`SELECT file_path FROM uploads WHERE session_id IS NOT NULL AND created_at < NOW() - INTERVAL 1 HOUR`)
	if err != nil {
		return utils.ErrorHandler(err, "Unable to find the file path.")
	}

	defer rows.Close()

	for rows.Next() {
		var path string

		rows.Scan(&path)

		err = os.Remove(`C:\Users\Shashank.BR\OneDrive\Desktop\Go programing\Project\BLOG_Rest_API\cmd\server\` + path)
		if err != nil {
			return utils.ErrorHandler(err, "Unable to remove the files")
		}

		_, err := db.Exec(`DELETE FROM uploads WHERE file_path=?`, path)
		if err != nil {
			return utils.ErrorHandler(err, "Unable to delete the file path in DB")
		}

	}
	return nil
}

func SessionsCleanUp() error {
	db, err := db.ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM upload_sessions WHERE created_at < NOW() - INTERVAL 1 HOUR")
	if err != nil {
		return utils.ErrorHandler(err, "Unable to delete the session id")
	} else {
		log.Println("Old Sessions Cleaned")
	}
	return nil
}

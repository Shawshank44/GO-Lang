package sqlconnect

import (
	"context"
	"log"
	"order_mgt/pkg/storage"
	"order_mgt/pkg/utils"
)

func UploadCleanUp(ctx context.Context, minio *storage.MinioService) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	rows, err := db.QueryContext(ctx, `SELECT file_path FROM files WHERE session_id IS NOT NULL AND created_at < NOW() - INTERVAL 1 HOUR`)
	if err != nil {
		return utils.ErrorHandler(err, "unable to find expired files")
	}

	var paths []string

	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			rows.Close()
			return utils.ErrorHandler(err, "unable to scan the file path")
		}
		paths = append(paths, path)
	}

	if err := rows.Err(); err != nil {
		rows.Close()
		return utils.ErrorHandler(err, "error while reading the expired file")
	}

	rows.Close()

	for _, path := range paths {
		if err := minio.Delete(ctx, path); err != nil {
			log.Printf("cleanup: unable to delete MinIO object %s: %v", path, err)
			continue
		}

		_, err := db.ExecContext(ctx, `DELETE FROM files WHERE file_path = ?`, path)
		if err != nil {
			log.Printf("cleanup: MinIO object deleted but DB record could not be deleted %s: %v", path, err)
			continue
		}
		log.Printf("cleanup: deleted unused file %s", path)
	}

	_, err = db.ExecContext(ctx, `DELETE FROM upload_sessions WHERE created_at < NOW() - INTERVAL 1 HOUR AND NOT EXISTS (SELECT 1 FROM files WHERE files.session_id = upload_sessions.id)`)
	if err != nil {
		return utils.ErrorHandler(err, "unable to delete expired upload sessions")
	}

	log.Println("Upload cleanup completed")

	return nil
}

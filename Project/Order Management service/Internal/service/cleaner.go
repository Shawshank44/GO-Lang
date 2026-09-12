package service

import (
	"context"
	"log"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	"order_mgt/pkg/storage"
	"order_mgt/pkg/utils"
	"time"
)

func StartCleaner(ctx context.Context, minio *storage.MinioService) error {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Cleanup running...")
			if err := sqlconnect.UploadCleanUp(ctx, minio); err != nil {
				return utils.ErrorHandler(err, "Unable to run the cleaner")
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

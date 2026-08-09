package utilssql

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"order_mgt/pkg/utils"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioService struct {
	Client     *minio.Client
	BucketName string
}

func NewMinioService() (*MinioService, error) {
	port := fmt.Sprintf("localhost%s", os.Getenv("MINIO_PORT"))
	client, err := minio.New(port, &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})

	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to start the MINIO service")
	}

	return &MinioService{
		Client:     client,
		BucketName: "product-images",
	}, nil
}

func (m *MinioService) CheckBucket() error {
	ctx := context.Background()

	exists, err := m.Client.BucketExists(ctx, m.BucketName)
	if err != nil {
		return utils.ErrorHandler(err, "Invalid Bucket")
	}

	if !exists {
		return utils.ErrorHandler(errors.New("bucket does not exist"), "Bucket does not exist")
	}

	return nil
}

func (m *MinioService) Upload(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	objectName := fileHeader.Filename

	_, err := m.Client.PutObject(ctx, m.BucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})

	if err != nil {
		return "", utils.ErrorHandler(err, "Unable to upload the file")
	}

	return objectName, nil
}

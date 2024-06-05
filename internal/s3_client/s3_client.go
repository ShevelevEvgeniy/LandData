package s3_client

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

type MinioClient interface {
	UploadFile(ctx context.Context, fileName string, file io.Reader, size int64) error
	DownloadFile(fileName string) ([]byte, error)
	GetMinio() *minio.Client
	GetBucket() string
}

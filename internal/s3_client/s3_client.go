package s3_client

import (
	"context"
	"io"
	"net/url"

	"github.com/minio/minio-go/v7"
)

type MinioClient interface {
	UploadFile(ctx context.Context, fileName string, file io.Reader, size int64) error
	GetLinkDownload(ctx context.Context, fileName string) (*url.URL, error)
	GetMinio() *minio.Client
	GetBucket() string
}

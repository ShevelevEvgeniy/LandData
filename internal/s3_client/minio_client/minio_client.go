package minio_client

import (
	"context"
	"io"

	"github.com/ShevelevEvgeniy/app/config"
	def "github.com/ShevelevEvgeniy/app/internal/s3_client"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

var _ def.MinioClient = (*MinioClient)(nil)

type MinioClient struct {
	Bucket string
	client *minio.Client
}

func NewMinioClient(ctx context.Context, cfg *config.Config) (*MinioClient, error) {
	var err error
	var minioClient = &MinioClient{}

	minioClient.client, err = minio.New(
		cfg.Minio.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
			Secure: cfg.Minio.UseSSL,
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create minio_client s3_client")
	}

	minioClient.Bucket = cfg.Minio.BucketName

	err = minioClient.CreateBucket(ctx, minioClient.Bucket, cfg.Minio.Region)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func (m *MinioClient) CreateBucket(ctx context.Context, bucketName string, region string) error {
	exist, errBucketExists := m.client.BucketExists(ctx, bucketName)
	if exist && errBucketExists == nil {
		return nil
	}

	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: region})
	if err != nil {
		return errors.Wrap(err, "failed to create minio_client bucket")
	}

	return nil
}

func (m *MinioClient) UploadFile(ctx context.Context, fileName string, file io.Reader, size int64) error {
	_, err := m.client.PutObject(ctx, m.Bucket, fileName, file, size, minio.PutObjectOptions{})
	return err
}

func (m *MinioClient) DownloadFile(fileName string) ([]byte, error) {
	return nil, nil
}

func (m *MinioClient) GetMinio() *minio.Client {
	return m.client
}

func (m *MinioClient) GetBucket() string {
	return m.Bucket
}

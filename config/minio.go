package config

type Minio struct {
	Endpoint        string `envconfig:"MINIO_ENDPOINT" env-required:"true"`
	AccessKeyID     string `envconfig:"MINIO_ACCESS_KEY" env-required:"true"`
	SecretAccessKey string `envconfig:"MINIO_SECRET_KEY" env-required:"true"`
	UseSSL          bool   `envconfig:"MINIO_USE_SSL" env-default:"false"`
	Region          string `envconfig:"MINIO_REGION" env-default:"us-east-1"`
	BucketName      string `envconfig:"MINIO_BUCKET" env-required:"true"`
}

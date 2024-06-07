package config

import (
	"log/slog"
	"os"

	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Mode        *ProjectMode
	HTTPServer  *HTTPServer
	DB          *DB
	Auth        *Auth
	Minio       *Minio
	RetryConfig *RetryConfig
	IpInfo      *IpInfo
}

func MustLoad(log *slog.Logger) *Config {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		log.Error("Failed to load config: ", sl.Err(err))
		os.Exit(1)
	}

	log.Info("Loaded config: ",
		"Mode.Env", cfg.Mode.Env,
		"HTTPServer.Port", cfg.HTTPServer.Port,
		"HTTPServer.LocalPort", cfg.HTTPServer.LocalPort,
		"HTTPServer.Timeout", cfg.HTTPServer.Timeout,
		"HTTPServer.IdleTimeout", cfg.HTTPServer.IdleTimeout,
		"HTTPServer.StopTimeout", cfg.HTTPServer.StopTimeout,
		"DB.Host", cfg.DB.Host,
		"DB.Port", cfg.DB.Port,
		"DB.DBName", cfg.DB.DBName,
		"DB.SslMode", cfg.DB.SslMode,
		"DB.DriverName", cfg.DB.DriverName,
		"DB.MigrationUrl", cfg.DB.MigrationUrl,
		"Minio.Endpoint", cfg.Minio.Endpoint,
		"Minio.UseSSL", cfg.Minio.UseSSL,
		"Minio.Region", cfg.Minio.Region,
		"Minio.BucketName", cfg.Minio.BucketName,
		"RetryConfig.Attempts", cfg.RetryConfig.Attempts,
		"RetryConfig.Delay", cfg.RetryConfig.Delay,
	)

	return &cfg
}

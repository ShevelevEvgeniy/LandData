package config

import (
	"log/slog"
	"os"

	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	"github.com/joho/godotenv"
)

type ProjectMode struct {
	Env string `envconfig:"ENV" env-default:"development"`
}

func LoadProjectMode() string {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Failed to load .env file: ", sl.Err(err))
	}

	return os.Getenv("ENV")
}

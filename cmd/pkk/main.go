package main

import (
	"context"
	"os"

	"github.com/ShevelevEvgeniy/app/config"
	application "github.com/ShevelevEvgeniy/app/internal/app"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	log := sl.SetupLogger(config.LoadProjectMode())

	app := application.NewApp(log)

	if err := app.Run(ctx); err != nil {
		log.Error("Failed to start app: ", sl.Err(err))
		os.Exit(1)
	}
}

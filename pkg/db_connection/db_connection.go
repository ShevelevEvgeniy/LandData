package db_connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ShevelevEvgeniy/app/config"
	"github.com/pkg/errors"
)

func Connect(ctx context.Context, cfg *config.DB) (*pgxpool.Pool, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	urlExample := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		cfg.DriverName, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	connConfig, err := pgxpool.ParseConfig(urlExample)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse db url")
	}

	connConfig.MaxConns = cfg.MaxConns

	conn, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to db")
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "Failed to ping db")
	}

	return conn, nil
}

func Close(conn *pgxpool.Pool) error {
	if conn != nil {
		conn.Close()
		return nil
	}

	return nil
}

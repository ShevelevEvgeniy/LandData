package kpt_repository

import (
	"context"
	errs "errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"

	"github.com/ShevelevEvgeniy/app/internal/repository/model"
	repositoryQuery "github.com/ShevelevEvgeniy/app/internal/repository/repository_query"
	generateQuery "github.com/ShevelevEvgeniy/app/lib/generate_query"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const TableName = "kpt"

type KptRepository struct {
	pool *pgxpool.Pool
}

func NewKptRepository(pool *pgxpool.Pool) *KptRepository {
	return &KptRepository{
		pool: pool,
	}
}

func (kp *KptRepository) SaveKptInfo(ctx context.Context, kptInfo model.Kpt) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	query, values, err := generateQuery.GenerateInsertQuery(kptInfo, TableName, true)
	if err != nil {
		return errors.Wrap(err, "Failed create query db")
	}

	conn, err := kp.pool.Acquire(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to acquire connection from pool")
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "Failed insert db")
	}

	return nil
}

func (kp *KptRepository) GetKptInfo(_ context.Context, cadQuarter string) (model.Kpt, error) {
	return model.Kpt{}, nil
}

func (kp *KptRepository) GetKptDateFormation(ctx context.Context, cadQuarter string) (time.Time, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	var dateFormation time.Time
	query := repositoryQuery.GetKptDateFormation

	conn, err := kp.pool.Acquire(ctx)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "Failed to acquire connection from pool")
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, query, "36:25:0000021").Scan(&dateFormation)
	if err != nil {
		if errs.Is(err, pgx.ErrNoRows) {
			return time.Time{}, nil
		}

		return time.Time{}, errors.Wrap(err, "Failed get date formation")
	}

	return dateFormation, nil
}

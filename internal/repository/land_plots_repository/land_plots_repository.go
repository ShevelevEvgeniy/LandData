package land_plots_repository

import (
	"context"
	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/internal/repository/model"
	repositoryQuery "github.com/ShevelevEvgeniy/app/internal/repository/repository_query"
	generateQuery "github.com/ShevelevEvgeniy/app/lib/generate_query"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

const tableName = "land_plots"

type LandPlotsRepository struct {
	pool *pgxpool.Pool
}

func NewLandPlotsRepository(pool *pgxpool.Pool) *LandPlotsRepository {
	return &LandPlotsRepository{
		pool: pool,
	}
}

func (lp *LandPlotsRepository) GetCoordinates(ctx context.Context, CadNumber string) (string, error) {
	var polygonText string

	conn, err := lp.pool.Acquire(ctx)
	if err != nil {
		return "", errors.Wrap(err, "Failed to acquire connection from pool")
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, repositoryQuery.GetCoordinatesByCadNumber, CadNumber).Scan(&polygonText)

	if err != nil {
		return "", errors.Wrap(err, "Failed get coordinates from db")
	}

	return polygonText, nil
}

func (lp *LandPlotsRepository) SaveLandPlots(ctx context.Context, landPlots []model.LandPlot) error {
	query, values, err := generateQuery.GenerateMultiInsertQuery(landPlots, tableName, true)
	if err != nil {
		return errors.Wrap(err, "Failed create query db")
	}

	conn, err := lp.pool.Acquire(ctx)
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

func (lp *LandPlotsRepository) GetLandPlots(ctx context.Context) ([]dto.LandPlot, error) {
	return nil, nil
}

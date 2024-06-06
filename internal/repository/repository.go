package repository

import (
	"context"
	"github.com/ShevelevEvgeniy/app/internal/dto"
	"time"

	"github.com/ShevelevEvgeniy/app/internal/repository/model"
)

type LandPlotsRepository interface {
	SaveLandPlots(ctx context.Context, LandPlots []model.LandPlot) error
	GetCoordinates(ctx context.Context, cadNumber string) (string, error)
}

type KptRepository interface {
	SaveKptInfo(ctx context.Context, kptInfo model.Kpt) error
	GetKptInfo(ctx context.Context, cadQuarter string) (dto.KptInfo, error)
	GetKptDateFormation(ctx context.Context, cadQuarter string) (time.Time, error)
	GetKptName(ctx context.Context, cadQuarter string) (string, error)
}

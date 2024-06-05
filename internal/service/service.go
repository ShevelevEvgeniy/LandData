package service

import (
	"context"

	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/internal/repository/model"
)

type LandPlotsService interface {
	GetCoordinatesList(ctx context.Context, cadNumber string) (dto.LandPlotCoordinates, error)
	SaveLandPlotsFromKpt(ctx context.Context, dto *dto.KptDto) error
}

type KptService interface {
	ExistKpt(ctx context.Context, kptInfo *model.Kpt) error
	SaveKptInfo(ctx context.Context, kptInfo *model.Kpt) error
	UploadFileToMinio(ctx context.Context, dto *dto.KptDto) error
}

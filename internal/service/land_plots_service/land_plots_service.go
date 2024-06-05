package land_plots_service

import (
	"context"

	"github.com/ShevelevEvgeniy/app/internal/converter"
	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/internal/repository"
	"github.com/paulmach/orb/encoding/wkt"
	"github.com/pkg/errors"
)

type LandPlotsService struct {
	repository repository.LandPlotsRepository
}

func NewLandPlotsService(repository repository.LandPlotsRepository) *LandPlotsService {
	return &LandPlotsService{
		repository: repository,
	}
}

func (lp *LandPlotsService) GetCoordinatesList(ctx context.Context, cadNumber string) (dto.LandPlotCoordinates, error) {
	var coordinates dto.LandPlotCoordinates

	polygonText, err := lp.repository.GetCoordinates(ctx, cadNumber)
	if err != nil {
		return coordinates, errors.Wrap(err, "failed to get coordinates")
	}

	polygon, err := wkt.Unmarshal(polygonText)
	if err != nil {
		return coordinates, errors.Wrap(err, "failed to unmarshal coordinates")
	}

	coordinates.CadNumber = cadNumber
	coordinates.Coordinates, err = converter.CoordinatesToPoints(polygon)
	if err != nil {
		return coordinates, errors.Wrap(err, "failed to convert coordinates")
	}

	return coordinates, nil
}

func (lp *LandPlotsService) SaveLandPlotsFromKpt(ctx context.Context, dto *dto.KptDto) error {
	return lp.repository.SaveLandPlots(ctx, converter.ToListLandPlotsFromKpt(dto.Territory))
}

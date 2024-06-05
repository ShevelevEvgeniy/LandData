package converter

import (
	"time"

	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/internal/repository/model"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/pkg/errors"
)

func ToKptInfoFromKpt(dto *dto.KptDto) (*model.Kpt, error) {
	const layout = "2006-01-02"

	data, err := time.Parse(layout, dto.Territory.DetailsStatement.GroupTopRequisites.DateFormation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse time")
	}

	return &model.Kpt{
		CadQuarter:    dto.Territory.CadastralBlocks[0].CadastralNumber,
		DateFormation: data,
		FileName:      dto.KptHeaders.Filename,
		Size:          dto.KptHeaders.Size,
		ContentType:   dto.KptHeaders.Header.Get("Content-Type"),
	}, nil
}

func ToListLandPlotsFromKpt(territory *dto.ExtractCadastralPlanTerritory) []model.LandPlot {
	const objectType = "Земельный участок"

	var landPlots []model.LandPlot

	for _, block := range territory.CadastralBlocks {
		for _, record := range block.LandRecords {
			if len(record.ContoursLocation.Contours) > 0 && record.Object.CommonData.Type.Value == objectType {
				var coordinates orb.Polygon

				for _, contour := range record.ContoursLocation.Contours {
					var ring orb.Ring

					for _, spatialElement := range contour.EntitySpatial.SpatialsElements {
						for _, ordinate := range spatialElement.Ordinates {
							ring = append(ring, orb.Point{ordinate.X, ordinate.Y})
						}
					}

					if len(ring) > 0 && ring[0] != ring[len(ring)-1] {
						ring = append(ring, ring[0])
					}

					if len(ring) < 4 {
						continue
					}

					coordinates = append(coordinates, ring)
				}

				if len(coordinates) == 0 {
					continue
				}
				wkbGeometry, _ := wkb.Marshal(coordinates)

				landPlot := model.LandPlot{
					CadNumber:       record.Object.CommonData.CadNumber,
					Coordinates:     wkbGeometry,
					Category:        record.Params.Category.Type.Value,
					PermittedUse:    record.Params.PermittedUse.PermittedUseEstablished.ByDocument,
					Area:            record.Params.Area.Value,
					Okato:           record.AddressLocation.Address.AddressFIAS.LevelSettlement.Okato,
					Kladr:           record.AddressLocation.Address.AddressFIAS.LevelSettlement.Kladr,
					ReadableAddress: record.AddressLocation.Address.ReadableAddress,
				}

				landPlots = append(landPlots, landPlot)
			}
		}
	}

	return landPlots
}

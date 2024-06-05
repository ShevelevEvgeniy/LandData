package converter

import (
	"fmt"
	"strconv"

	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/paulmach/orb"
)

func CoordinatesToPoints(geometry orb.Geometry) ([]dto.Point, error) {
	var points []dto.Point
	var pointNumber int

	switch g := geometry.(type) {
	case orb.Polygon:
		polygon := g
		for _, ring := range polygon {
			for _, point := range ring {
				points = append(points, dto.Point{
					Number: strconv.Itoa((pointNumber % (len(ring) - 1)) + 1),
					X:      point[0],
					Y:      point[1],
				})
				pointNumber++
			}
		}
	default:
		return nil, fmt.Errorf("unsupported geometry type")
	}

	return points, nil
}

package land_plots_handler

import (
	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/lib/api/response"
)

type Response struct {
	Status      response.Response
	CadNumber   string      `json:"cad_number,omitempty"`
	Coordinates []dto.Point `json:"coordinates,omitempty"`
}

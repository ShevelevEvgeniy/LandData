package dto

import "github.com/paulmach/orb"

type LandPlot struct {
	ID              string       `json:"id"`
	CadNumber       string       `json:"cad_number" validate:"required,cad_number"`
	Coordinates     orb.Geometry `json:"coordinates"`
	Category        string       `json:"category"`
	PermittedUse    string       `json:"permitted_use"`
	Area            float64      `json:"area"`
	Okato           string       `json:"okato"`
	Kladr           string       `json:"kladr"`
	ReadableAddress string       `json:"readable_address"`
}

package dto

import "time"

type KptInfo struct {
	CadQuarter      string    `json:"cad_quarter"`
	DateFormation   time.Time `json:"date_formation"`
	Link            string    `json:"link"`
	AmountLandPlots int       `json:"amount_land_plots"`
}

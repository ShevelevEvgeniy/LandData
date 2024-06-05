package repository_query

import (
	_ "embed"
)

var (
	//go:embed land_plots/get_coordinates_by_cad_number.sql
	GetCoordinatesByCadNumber string

	//go:embed layout_query/save_or_update.sql
	SaveOrUpdate string

	//go:embed layout_query/save.sql
	Save string

	//go:embed kpt/get_kpt_date_formation.sql
	GetKptDateFormation string
)

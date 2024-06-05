package dto

type Point struct {
	Number string  `json:"point_number"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}

type LandPlotCoordinates struct {
	CadNumber   string  `json:"cad_number"`
	Coordinates []Point `json:"coordinates"`
}

package model

type LandPlot struct {
	ID              string  `json:"id"`
	CadNumber       string  `json:"cad_number" unique:"true"`
	Coordinates     []byte  `json:"coordinates" description:"ST_GeomFromWKB"`
	Category        string  `json:"category"`
	PermittedUse    string  `json:"permitted_use"`
	Area            float64 `json:"area"`
	Okato           string  `json:"okato"`
	Kladr           string  `json:"kladr"`
	ReadableAddress string  `json:"readable_address"`
}

package model

import "time"

type Kpt struct {
	ID            string    `json:"id"`
	CadQuarter    string    `json:"cad_quarter" unique:"true"`
	DateFormation time.Time `json:"date_formation"`
	FileName      string    `json:"file_name"`
	Size          int64     `json:"size"`
	ContentType   string    `json:"content_type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

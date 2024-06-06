package dto

import (
	"bytes"
	"encoding/xml"
	"mime/multipart"
)

type KptDto struct {
	CadQuarter string `json:"cad_quarter" validate:"required,cad_quarter"`
	Kpt        bytes.Buffer
	KptHeaders *multipart.FileHeader
	Territory  *ExtractCadastralPlanTerritory
}

type ExtractCadastralPlanTerritory struct {
	XMLName          xml.Name         `xml:"extract_cadastral_plan_territory"`
	DetailsStatement DetailsStatement `xml:"details_statement"`
	DetailsRequest   DetailsRequest   `xml:"details_request"`
	CadastralBlocks  []CadastralBlock `xml:"cadastral_blocks>cadastral_block"`
}

type DetailsStatement struct {
	GroupTopRequisites GroupTopRequisites `xml:"group_top_requisites"`
}

type GroupTopRequisites struct {
	OrganRegistrRights string `xml:"organ_registr_rights"`
	DateFormation      string `xml:"date_formation"`
	RegistrationNumber string `xml:"registration_number"`
}

type DetailsRequest struct {
	DateReceivedRequest                  string `xml:"date_received_request"`
	DateReceiptRequestRegAuthorityRights string `xml:"date_receipt_request_reg_authority_rights"`
}

type CadastralBlock struct {
	CadastralNumber string       `xml:"cadastral_number"`
	LandRecords     []LandRecord `xml:"record_data>base_data>land_records>land_record"`
}

type LandRecord struct {
	Object           Object           `xml:"object"`
	Params           Params           `xml:"params"`
	AddressLocation  AddressLocation  `xml:"address_location"`
	Cost             Cost             `xml:"cost"`
	ContoursLocation ContoursLocation `xml:"contours_location"`
}

type Object struct {
	CommonData CommonData `xml:"common_data"`
	Subtype    Subtype    `xml:"subtype"`
}

type CommonData struct {
	Type      Type   `xml:"type"`
	CadNumber string `xml:"cad_number"`
}

type Type struct {
	Code  string `xml:"code"`
	Value string `xml:"value"`
}

type Subtype struct {
	Code  string `xml:"code"`
	Value string `xml:"value"`
}

type Params struct {
	Category     Category     `xml:"category"`
	PermittedUse PermittedUse `xml:"permitted_use"`
	Area         Area         `xml:"area"`
}

type Category struct {
	Type Type `xml:"type"`
}

type PermittedUse struct {
	PermittedUseEstablished PermittedUseEstablished `xml:"permitted_use_established"`
}

type PermittedUseEstablished struct {
	ByDocument string `xml:"by_document"`
}

type Area struct {
	Value      float64 `xml:"value"`
	Inaccuracy float64 `xml:"inaccuracy"`
}

type AddressLocation struct {
	Address Address `xml:"address"`
}

type Address struct {
	AddressFIAS     AddressFIAS `xml:"address_fias"`
	ReadableAddress string      `xml:"readable_address"`
}

type AddressFIAS struct {
	LevelSettlement LevelSettlement `xml:"level_settlement"`
	DetailedLevel   DetailedLevel   `xml:"detailed_level"`
}

type LevelSettlement struct {
	Okato    string   `xml:"okato"`
	Kladr    string   `xml:"kladr"`
	Region   Region   `xml:"region"`
	City     City     `xml:"city"`
	Locality Locality `xml:"locality"`
}

type Region struct {
	Code  string `xml:"code"`
	Value string `xml:"value"`
}

type City struct {
	TypeCity string `xml:"type_city"`
	NameCity string `xml:"name_city"`
}

type Locality struct {
	TypeLocality string `xml:"type_locality"`
	NameLocality string `xml:"name_locality"`
}

type DetailedLevel struct {
	Street Street `xml:"street"`
	Level1 Level1 `xml:"level1"`
}

type Street struct {
	TypeStreet string `xml:"type_street"`
	NameStreet string `xml:"name_street"`
}

type Level1 struct {
	TypeLevel1 string `xml:"type_level1"`
	NameLevel1 string `xml:"name_level1"`
}

type Cost struct {
	Value float64 `xml:"value"`
}

type ContoursLocation struct {
	Contours []Contour `xml:"contours>contour"`
}

type Contour struct {
	EntitySpatial EntitySpatial `xml:"entity_spatial"`
}

type EntitySpatial struct {
	SkID             string           `xml:"sk_id"`
	SpatialsElements []SpatialElement `xml:"spatials_elements>spatial_element"`
}

type SpatialElement struct {
	Ordinates []Ordinate `xml:"ordinates>ordinate"`
}

type Ordinate struct {
	X             float64 `xml:"x"`
	Y             float64 `xml:"y"`
	OrdNmb        int     `xml:"ord_nmb"`
	NumGeoPoint   string  `xml:"num_geopoint"`
	DeltaGeoPoint float64 `xml:"delta_geopoint"`
}

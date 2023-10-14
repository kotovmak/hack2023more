package model

type TypeListRequest struct {
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Radius     float64  `json:"radius"`
	Tags       []string `json:"tags"`
	TagsUslugi []string `json:"tags_uslugi"`
}

type TypeList struct {
	TotalAtms    int      `json:"total_atms"`
	TotalOffices int      `json:"total_offices"`
	Atm          []Atm    `json:"atms"`
	Office       []Office `json:"offices"`
}

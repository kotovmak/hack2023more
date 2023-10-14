package model

type Office struct {
	ID                  int     `json:"id"`
	Load                string  `json:"load"`
	Rating              float64 `json:"rating"`
	SalePointName       string  `json:"salePointName"`
	Address             string  `json:"address"`
	Status              string  `json:"status"`
	OpenHours           Days    `json:"openHours"`
	Rko                 string  `json:"rko"`
	OpenHoursIndividual Days    `json:"openHoursIndividual"`
	OfficeType          string  `json:"officeType"`
	SalePointFormat     string  `json:"salePointFormat"`
	SuoAvailability     string  `json:"suoAvailability"`
	HasRamp             string  `json:"hasRamp"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	MetroStation        any     `json:"metroStation"`
	Distance            int     `json:"distance"`
	Kep                 bool    `json:"kep"`
	MyBranch            bool    `json:"myBranch"`
	Network             any     `json:"network,omitempty"`
	SalePointCode       string  `json:"salePointCode,omitempty"`
}

type Days struct {
	Day0 Day `json:"0"`
	Day1 Day `json:"1"`
	Day2 Day `json:"2"`
	Day3 Day `json:"3"`
	Day4 Day `json:"4"`
	Day5 Day `json:"5"`
	Day6 Day `json:"6"`
}

type Day struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}

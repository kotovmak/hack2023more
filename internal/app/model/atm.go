package model

type Atm struct {
	ID        int     `json:"id"`
	Load      string  `json:"load"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	AllDay    bool    `json:"allDay"`
	Services  struct {
		Wheelchair struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"wheelchair"`
		Blind struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"blind"`
		NfcForBankCards struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"nfcForBankCards"`
		QrRead struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"qrRead"`
		SupportsUsd struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"supportsUsd"`
		SupportsChargeRub struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"supportsChargeRub"`
		SupportsEur struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"supportsEur"`
		SupportsRub struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"supportsRub"`
	} `json:"services"`
}

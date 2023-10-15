package server

import (
	"database/sql"
	"hack2023/internal/app/model"
	"log"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) typelist(c echo.Context) error {
	a, err := s.store.ListAtms(c.Request().Context())
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
		return echo.ErrInternalServerError
	}
	o, err := s.store.ListOffices(c.Request().Context())
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
		return echo.ErrInternalServerError
	}
	tl := model.TypeList{
		TotalAtms:    len(a),
		TotalOffices: len(o),
		Atm:          a,
		Office:       o,
	}

	return c.JSON(http.StatusOK, tl)
}

func (s *server) typelistFilter(c echo.Context) error {
	r := new(model.TypeListRequest)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if r.Radius == 0 {
		r.Radius = 5
	}

	a, err := s.store.ListAtms(c.Request().Context())
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
		return echo.ErrInternalServerError
	}
	o, err := s.store.ListOffices(c.Request().Context())
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
		return echo.ErrInternalServerError
	}

	minO := []model.Office{}
	for _, v := range o {
		ok := true
		distance := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(deg2rad((r.Latitude-v.Latitude)/2)), 2)+
			math.Cos(deg2rad(r.Latitude))*math.Cos(deg2rad(v.Latitude))*
				math.Pow(math.Sin(deg2rad((r.Longitude-v.Longitude)/2)), 2))) * 6378245
		if distance > (r.Radius * 1000) {
			ok = false
		}
		for i := range r.Tags {
			if r.Tags[i] == "wheelchair" && v.HasRamp != "Y" {
				ok = false
			}
		}
		if ok {
			v.Load = "middle"
			minO = append(minO, v)
		}
	}
	minA := []model.Atm{}
	for _, v := range a {
		ok := true
		distance := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(deg2rad((r.Latitude-v.Latitude)/2)), 2)+
			math.Cos(deg2rad(r.Latitude))*math.Cos(deg2rad(v.Latitude))*
				math.Pow(math.Sin(deg2rad((r.Longitude-v.Longitude)/2)), 2))) * 6378245
		if distance > (r.Radius * 1000) {
			ok = false
		}
		for i := range r.Tags {
			if r.Tags[i] == "wheelchair" && v.Services.Wheelchair.ServiceActivity == "UNAVAILABLE" {
				ok = false
			}
			if r.Tags[i] == "blind" && v.Services.Blind.ServiceActivity == "UNAVAILABLE" {
				ok = false
			}
			if r.Tags[i] == "nfcForBankCards" && v.Services.NfcForBankCards.ServiceActivity == "UNAVAILABLE" {
				ok = false
			}
			if r.Tags[i] == "qrRead" && v.Services.QrRead.ServiceActivity == "UNAVAILABLE" {
				ok = false
			}
			if r.Tags[i] == "supportsUsd" && v.Services.SupportsUsd.ServiceActivity == "UNAVAILABLE" {
				ok = false
			}
			if r.Tags[i] == "supportsRub" && v.Services.SupportsRub.ServiceActivity == "UNAVAILABLE" {
				ok = false
			}
			if r.Tags[i] == "supportsEur" && v.Services.SupportsEur.ServiceActivity == "UNAVAILABLE" {
				ok = false
			}
			if r.Tags[i] == "supportsChargeRub" && v.Services.SupportsChargeRub.ServiceActivity == "UNAVAILABLE" {
				ok = false
			}
		}
		if ok {
			v.Load = "low"
			minA = append(minA, v)
		}
	}

	if len(minO) > 0 {
		minO, err = s.RunOfficePredict(c.Request().Context(), minO)
		if err != nil {
			log.Print(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	tl := model.TypeList{
		TotalAtms:    len(minA),
		TotalOffices: len(minO),
		Atm:          minA,
		Office:       minO,
	}

	return c.JSON(http.StatusOK, tl)
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

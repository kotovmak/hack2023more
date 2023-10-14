package store

import (
	"context"
	"database/sql"
	"hack2023/internal/app/model"
)

const createAtm = `-- name: CreateAtm :execresult
INSERT INTO atms(
  address,
  latitude,
  longitude,
  allDay,
  serviceswheelchairserviceCapability,
  serviceswheelchairserviceActivity,
  servicesblindserviceCapability,
  servicesblindserviceActivity,
  servicesnfcForBankCardsserviceCapability,
  servicesnfcForBankCardsserviceActivity,
  servicesqrReadserviceCapability,
  servicesqrReadserviceActivity,servicessupportsUsdserviceCapability,servicessupportsUsdserviceActivity,servicessupportsChargeRubserviceCapability,servicessupportsChargeRubserviceActivity,servicessupportsEurserviceCapability,servicessupportsEurserviceActivity,servicessupportsRubserviceCapability,servicessupportsRubserviceActivity) 
  VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

type CreateAtmParams struct {
	Address                                    string
	Latitude                                   string
	Longitude                                  string
	Allday                                     bool
	Serviceswheelchairservicecapability        string
	Serviceswheelchairserviceactivity          string
	Servicesblindservicecapability             string
	Servicesblindserviceactivity               string
	Servicesnfcforbankcardsservicecapability   string
	Servicesnfcforbankcardsserviceactivity     string
	Servicesqrreadservicecapability            string
	Servicesqrreadserviceactivity              string
	Servicessupportsusdservicecapability       string
	Servicessupportsusdserviceactivity         string
	Servicessupportschargerubservicecapability string
	Servicessupportschargerubserviceactivity   string
	Servicessupportseurservicecapability       string
	Servicessupportseurserviceactivity         string
	Servicessupportsrubservicecapability       string
	Servicessupportsrubserviceactivity         string
}

func (q *Store) CreateAtm(ctx context.Context, arg CreateAtmParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAtm,
		arg.Address,
		arg.Latitude,
		arg.Longitude,
		arg.Allday,
		arg.Serviceswheelchairservicecapability,
		arg.Serviceswheelchairserviceactivity,
		arg.Servicesblindservicecapability,
		arg.Servicesblindserviceactivity,
		arg.Servicesnfcforbankcardsservicecapability,
		arg.Servicesnfcforbankcardsserviceactivity,
		arg.Servicesqrreadservicecapability,
		arg.Servicesqrreadserviceactivity,
		arg.Servicessupportsusdservicecapability,
		arg.Servicessupportsusdserviceactivity,
		arg.Servicessupportschargerubservicecapability,
		arg.Servicessupportschargerubserviceactivity,
		arg.Servicessupportseurservicecapability,
		arg.Servicessupportseurserviceactivity,
		arg.Servicessupportsrubservicecapability,
		arg.Servicessupportsrubserviceactivity,
	)
}

const deleteAtm = `-- name: DeleteAtm :exec
DELETE FROM atms
WHERE id = ?
`

func (q *Store) DeleteAtm(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAtm, id)
	return err
}

const getAtm = `-- name: GetAtm :one
SELECT id, address, latitude, longitude, allday, serviceswheelchairservicecapability, serviceswheelchairserviceactivity, servicesblindservicecapability, servicesblindserviceactivity, servicesnfcforbankcardsservicecapability, servicesnfcforbankcardsserviceactivity, servicesqrreadservicecapability, servicesqrreadserviceactivity, servicessupportsusdservicecapability, servicessupportsusdserviceactivity, servicessupportschargerubservicecapability, servicessupportschargerubserviceactivity, servicessupportseurservicecapability, servicessupportseurserviceactivity, servicessupportsrubservicecapability, servicessupportsrubserviceactivity FROM atms
WHERE id = ? LIMIT 1
`

func (q *Store) GetAtm(ctx context.Context, id int32) (model.Atm, error) {
	row := q.db.QueryRowContext(ctx, getAtm, id)
	var (
		i       model.Atm
		address sql.NullString
	)
	err := row.Scan(
		&i.ID,
		&address,
		&i.Latitude,
		&i.Longitude,
		&i.AllDay,
		&i.Services.Wheelchair.ServiceCapability,
		&i.Services.Wheelchair.ServiceActivity,
		&i.Services.Blind.ServiceCapability,
		&i.Services.Blind.ServiceActivity,
		&i.Services.NfcForBankCards.ServiceCapability,
		&i.Services.NfcForBankCards.ServiceActivity,
		&i.Services.QrRead.ServiceCapability,
		&i.Services.QrRead.ServiceActivity,
		&i.Services.SupportsUsd.ServiceCapability,
		&i.Services.SupportsUsd.ServiceActivity,
		&i.Services.SupportsChargeRub.ServiceCapability,
		&i.Services.SupportsChargeRub.ServiceActivity,
		&i.Services.SupportsEur.ServiceCapability,
		&i.Services.SupportsEur.ServiceActivity,
		&i.Services.SupportsRub.ServiceCapability,
		&i.Services.SupportsRub.ServiceActivity,
	)
	i.Address = address.String
	return i, err
}

const listAtms = `-- name: ListAtms :many
SELECT id, address, latitude, longitude, allday, serviceswheelchairservicecapability, serviceswheelchairserviceactivity, servicesblindservicecapability, servicesblindserviceactivity, servicesnfcforbankcardsservicecapability, servicesnfcforbankcardsserviceactivity, servicesqrreadservicecapability, servicesqrreadserviceactivity, servicessupportsusdservicecapability, servicessupportsusdserviceactivity, servicessupportschargerubservicecapability, servicessupportschargerubserviceactivity, servicessupportseurservicecapability, servicessupportseurserviceactivity, servicessupportsrubservicecapability, servicessupportsrubserviceactivity FROM atms
ORDER BY address
`

func (q *Store) ListAtms(ctx context.Context) ([]model.Atm, error) {
	rows, err := q.db.QueryContext(ctx, listAtms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.Atm
	for rows.Next() {
		var (
			i       model.Atm
			address sql.NullString
		)
		if err := rows.Scan(
			&i.ID,
			&address,
			&i.Latitude,
			&i.Longitude,
			&i.AllDay,
			&i.Services.Wheelchair.ServiceCapability,
			&i.Services.Wheelchair.ServiceActivity,
			&i.Services.Blind.ServiceCapability,
			&i.Services.Blind.ServiceActivity,
			&i.Services.NfcForBankCards.ServiceCapability,
			&i.Services.NfcForBankCards.ServiceActivity,
			&i.Services.QrRead.ServiceCapability,
			&i.Services.QrRead.ServiceActivity,
			&i.Services.SupportsUsd.ServiceCapability,
			&i.Services.SupportsUsd.ServiceActivity,
			&i.Services.SupportsChargeRub.ServiceCapability,
			&i.Services.SupportsChargeRub.ServiceActivity,
			&i.Services.SupportsEur.ServiceCapability,
			&i.Services.SupportsEur.ServiceActivity,
			&i.Services.SupportsRub.ServiceCapability,
			&i.Services.SupportsRub.ServiceActivity,
		); err != nil {
			return nil, err
		}
		i.Address = address.String
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

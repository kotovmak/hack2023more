package store

import (
	"context"
	"database/sql"
	"hack2023/internal/app/model"
)

const createOffice = `-- name: CreateOffice :execresult
INSERT INTO offices(
  salePointName,address,status,openHours0days,openHours0hours,openHours1days,openHours1hours,openHours2days,openHours2hours,openHours3days,openHours3hours,openHours4days,openHours4hours,openHours5days,openHours5hours,openHours6days,openHours6hours,rko,openHoursIndividual0days,openHoursIndividual0hours,openHoursIndividual1days,openHoursIndividual1hours,openHoursIndividual2days,openHoursIndividual2hours,openHoursIndividual3days,openHoursIndividual3hours,openHoursIndividual4days,openHoursIndividual4hours,openHoursIndividual5days,openHoursIndividual5hours,openHoursIndividual6days,openHoursIndividual6hours,officeType,salePointFormat,suoAvailability,hasRamp,latitude,longitude,metroStation,distance,kep,myBranch,network,salePointCode) 
  VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`

type CreateOfficeParams struct {
	Salepointname             string
	Address                   string
	Status                    string
	Openhours0days            string
	Openhours0hours           sql.NullString
	Openhours1days            sql.NullString
	Openhours1hours           sql.NullString
	Openhours2days            sql.NullString
	Openhours2hours           sql.NullString
	Openhours3days            sql.NullString
	Openhours3hours           sql.NullString
	Openhours4days            sql.NullString
	Openhours4hours           sql.NullString
	Openhours5days            sql.NullString
	Openhours5hours           sql.NullString
	Openhours6days            sql.NullString
	Openhours6hours           sql.NullString
	Rko                       sql.NullString
	Openhoursindividual0days  string
	Openhoursindividual0hours sql.NullString
	Openhoursindividual1days  sql.NullString
	Openhoursindividual1hours sql.NullString
	Openhoursindividual2days  sql.NullString
	Openhoursindividual2hours sql.NullString
	Openhoursindividual3days  sql.NullString
	Openhoursindividual3hours sql.NullString
	Openhoursindividual4days  sql.NullString
	Openhoursindividual4hours sql.NullString
	Openhoursindividual5days  sql.NullString
	Openhoursindividual5hours sql.NullString
	Openhoursindividual6days  sql.NullString
	Openhoursindividual6hours sql.NullString
	Officetype                string
	Salepointformat           string
	Suoavailability           sql.NullString
	Hasramp                   sql.NullString
	Latitude                  string
	Longitude                 string
	Metrostation              sql.NullString
	Distance                  int32
	Kep                       sql.NullBool
	Mybranch                  bool
	Network                   sql.NullString
	Salepointcode             sql.NullString
}

func (q *Store) CreateOffice(ctx context.Context, arg CreateOfficeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOffice,
		arg.Salepointname,
		arg.Address,
		arg.Status,
		arg.Openhours0days,
		arg.Openhours0hours,
		arg.Openhours1days,
		arg.Openhours1hours,
		arg.Openhours2days,
		arg.Openhours2hours,
		arg.Openhours3days,
		arg.Openhours3hours,
		arg.Openhours4days,
		arg.Openhours4hours,
		arg.Openhours5days,
		arg.Openhours5hours,
		arg.Openhours6days,
		arg.Openhours6hours,
		arg.Rko,
		arg.Openhoursindividual0days,
		arg.Openhoursindividual0hours,
		arg.Openhoursindividual1days,
		arg.Openhoursindividual1hours,
		arg.Openhoursindividual2days,
		arg.Openhoursindividual2hours,
		arg.Openhoursindividual3days,
		arg.Openhoursindividual3hours,
		arg.Openhoursindividual4days,
		arg.Openhoursindividual4hours,
		arg.Openhoursindividual5days,
		arg.Openhoursindividual5hours,
		arg.Openhoursindividual6days,
		arg.Openhoursindividual6hours,
		arg.Officetype,
		arg.Salepointformat,
		arg.Suoavailability,
		arg.Hasramp,
		arg.Latitude,
		arg.Longitude,
		arg.Metrostation,
		arg.Distance,
		arg.Kep,
		arg.Mybranch,
		arg.Network,
		arg.Salepointcode,
	)
}

const deleteOffice = `-- name: DeleteOffice :exec
DELETE FROM offices
WHERE id = ?
`

func (q *Store) DeleteOffice(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteOffice, id)
	return err
}

const getOffice = `-- name: GetOffice :one
SELECT id, salepointname, address, status, openhours0days, openhours0hours, openhours1days, openhours1hours, openhours2days, openhours2hours, openhours3days, openhours3hours, openhours4days, openhours4hours, openhours5days, openhours5hours, openhours6days, openhours6hours, rko, openhoursindividual0days, openhoursindividual0hours, openhoursindividual1days, openhoursindividual1hours, openhoursindividual2days, openhoursindividual2hours, openhoursindividual3days, openhoursindividual3hours, openhoursindividual4days, openhoursindividual4hours, openhoursindividual5days, openhoursindividual5hours, openhoursindividual6days, openhoursindividual6hours, officetype, salepointformat, suoavailability, hasramp, latitude, longitude, metrostation, distance, kep, mybranch, network, salepointcode FROM offices
WHERE id = ? LIMIT 1
`

func (q *Store) GetOffice(ctx context.Context, id int32) (model.Office, error) {
	row := q.db.QueryRowContext(ctx, getOffice, id)
	var (
		i                         model.Office
		openhours0days            sql.NullString
		openhours0hours           sql.NullString
		openhours1days            sql.NullString
		openhours1hours           sql.NullString
		openhours2days            sql.NullString
		openhours2hours           sql.NullString
		openhours3days            sql.NullString
		openhours3hours           sql.NullString
		openhours4days            sql.NullString
		openhours4hours           sql.NullString
		openhours5days            sql.NullString
		openhours5hours           sql.NullString
		openhours6days            sql.NullString
		openhours6hours           sql.NullString
		openhoursindividual0days  sql.NullString
		openhoursindividual0hours sql.NullString
		openhoursindividual1days  sql.NullString
		openhoursindividual1hours sql.NullString
		openhoursindividual2days  sql.NullString
		openhoursindividual2hours sql.NullString
		openhoursindividual3days  sql.NullString
		openhoursindividual3hours sql.NullString
		openhoursindividual4days  sql.NullString
		openhoursindividual4hours sql.NullString
		openhoursindividual5days  sql.NullString
		openhoursindividual5hours sql.NullString
		openhoursindividual6days  sql.NullString
		openhoursindividual6hours sql.NullString
		salepointcode             sql.NullString
		rko                       sql.NullString
		kep                       sql.NullBool
		suoavailability           sql.NullString
		hasramp                   sql.NullString
	)
	err := row.Scan(
		&i.ID,
		&i.SalePointName,
		&i.Address,
		&i.Status,
		&openhours0days,
		&openhours0hours,
		&openhours1days,
		&openhours1hours,
		&openhours2days,
		&openhours2hours,
		&openhours3days,
		&openhours3hours,
		&openhours4days,
		&openhours4hours,
		&openhours5days,
		&openhours5hours,
		&openhours6days,
		&openhours6hours,
		&rko,
		&openhoursindividual0days,
		&openhoursindividual0hours,
		&openhoursindividual1days,
		&openhoursindividual1hours,
		&openhoursindividual2days,
		&openhoursindividual2hours,
		&openhoursindividual3days,
		&openhoursindividual3hours,
		&openhoursindividual4days,
		&openhoursindividual4hours,
		&openhoursindividual5days,
		&openhoursindividual5hours,
		&openhoursindividual6days,
		&openhoursindividual6hours,
		&i.OfficeType,
		&i.SalePointFormat,
		&suoavailability,
		&hasramp,
		&i.Latitude,
		&i.Longitude,
		&i.MetroStation,
		&i.Distance,
		&kep,
		&i.MyBranch,
		&i.Network,
		&salepointcode,
	)
	i.OpenHours.Day0.Days = openhours0days.String
	i.OpenHours.Day0.Hours = openhours0hours.String
	i.OpenHours.Day1.Days = openhours1days.String
	i.OpenHours.Day1.Hours = openhours1hours.String
	i.OpenHours.Day2.Days = openhours2days.String
	i.OpenHours.Day2.Hours = openhours2hours.String
	i.OpenHours.Day3.Days = openhours3days.String
	i.OpenHours.Day3.Hours = openhours3hours.String
	i.OpenHours.Day4.Days = openhours4days.String
	i.OpenHours.Day4.Hours = openhours4hours.String
	i.OpenHours.Day5.Days = openhours5days.String
	i.OpenHours.Day5.Hours = openhours5hours.String
	i.OpenHours.Day6.Days = openhours6days.String
	i.OpenHours.Day6.Hours = openhours6hours.String
	i.OpenHoursIndividual.Day0.Days = openhoursindividual0days.String
	i.OpenHoursIndividual.Day0.Hours = openhoursindividual0hours.String
	i.OpenHoursIndividual.Day1.Days = openhoursindividual1days.String
	i.OpenHoursIndividual.Day1.Hours = openhoursindividual1hours.String
	i.OpenHoursIndividual.Day2.Days = openhoursindividual2days.String
	i.OpenHoursIndividual.Day2.Hours = openhoursindividual2hours.String
	i.OpenHoursIndividual.Day3.Days = openhoursindividual3days.String
	i.OpenHoursIndividual.Day3.Hours = openhoursindividual3hours.String
	i.OpenHoursIndividual.Day4.Days = openhoursindividual4days.String
	i.OpenHoursIndividual.Day4.Hours = openhoursindividual4hours.String
	i.OpenHoursIndividual.Day5.Days = openhoursindividual5days.String
	i.OpenHoursIndividual.Day5.Hours = openhoursindividual5hours.String
	i.OpenHoursIndividual.Day6.Days = openhoursindividual6days.String
	i.OpenHoursIndividual.Day6.Hours = openhoursindividual6hours.String
	i.SalePointCode = salepointcode.String
	i.Rko = rko.String
	i.Kep = kep.Bool
	i.SuoAvailability = suoavailability.String
	i.HasRamp = hasramp.String
	return i, err
}

const listOffices = `-- name: ListOffices :many
SELECT id, salepointname, address, status, openhours0days, openhours0hours, openhours1days, openhours1hours, openhours2days, openhours2hours, openhours3days, openhours3hours, openhours4days, openhours4hours, openhours5days, openhours5hours, openhours6days, openhours6hours, rko, openhoursindividual0days, openhoursindividual0hours, openhoursindividual1days, openhoursindividual1hours, openhoursindividual2days, openhoursindividual2hours, openhoursindividual3days, openhoursindividual3hours, openhoursindividual4days, openhoursindividual4hours, openhoursindividual5days, openhoursindividual5hours, openhoursindividual6days, openhoursindividual6hours, officetype, salepointformat, suoavailability, hasramp, latitude, longitude, metrostation, distance, kep, mybranch, network, salepointcode FROM offices
ORDER BY address
`

func (q *Store) ListOffices(ctx context.Context) ([]model.Office, error) {
	rows, err := q.db.QueryContext(ctx, listOffices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.Office
	for rows.Next() {
		var (
			i                         model.Office
			openhours0days            sql.NullString
			openhours0hours           sql.NullString
			openhours1days            sql.NullString
			openhours1hours           sql.NullString
			openhours2days            sql.NullString
			openhours2hours           sql.NullString
			openhours3days            sql.NullString
			openhours3hours           sql.NullString
			openhours4days            sql.NullString
			openhours4hours           sql.NullString
			openhours5days            sql.NullString
			openhours5hours           sql.NullString
			openhours6days            sql.NullString
			openhours6hours           sql.NullString
			openhoursindividual0days  sql.NullString
			openhoursindividual0hours sql.NullString
			openhoursindividual1days  sql.NullString
			openhoursindividual1hours sql.NullString
			openhoursindividual2days  sql.NullString
			openhoursindividual2hours sql.NullString
			openhoursindividual3days  sql.NullString
			openhoursindividual3hours sql.NullString
			openhoursindividual4days  sql.NullString
			openhoursindividual4hours sql.NullString
			openhoursindividual5days  sql.NullString
			openhoursindividual5hours sql.NullString
			openhoursindividual6days  sql.NullString
			openhoursindividual6hours sql.NullString
			salepointcode             sql.NullString
			rko                       sql.NullString
			kep                       sql.NullBool
			suoavailability           sql.NullString
			hasramp                   sql.NullString
		)

		if err := rows.Scan(
			&i.ID,
			&i.SalePointName,
			&i.Address,
			&i.Status,
			&openhours0days,
			&openhours0hours,
			&openhours1days,
			&openhours1hours,
			&openhours2days,
			&openhours2hours,
			&openhours3days,
			&openhours3hours,
			&openhours4days,
			&openhours4hours,
			&openhours5days,
			&openhours5hours,
			&openhours6days,
			&openhours6hours,
			&rko,
			&openhoursindividual0days,
			&openhoursindividual0hours,
			&openhoursindividual1days,
			&openhoursindividual1hours,
			&openhoursindividual2days,
			&openhoursindividual2hours,
			&openhoursindividual3days,
			&openhoursindividual3hours,
			&openhoursindividual4days,
			&openhoursindividual4hours,
			&openhoursindividual5days,
			&openhoursindividual5hours,
			&openhoursindividual6days,
			&openhoursindividual6hours,
			&i.OfficeType,
			&i.SalePointFormat,
			&suoavailability,
			&hasramp,
			&i.Latitude,
			&i.Longitude,
			&i.MetroStation,
			&i.Distance,
			&kep,
			&i.MyBranch,
			&i.Network,
			&salepointcode,
		); err != nil {
			return nil, err
		}
		i.OpenHours.Day0.Days = openhours0days.String
		i.OpenHours.Day0.Hours = openhours0hours.String
		i.OpenHours.Day1.Days = openhours1days.String
		i.OpenHours.Day1.Hours = openhours1hours.String
		i.OpenHours.Day2.Days = openhours2days.String
		i.OpenHours.Day2.Hours = openhours2hours.String
		i.OpenHours.Day3.Days = openhours3days.String
		i.OpenHours.Day3.Hours = openhours3hours.String
		i.OpenHours.Day4.Days = openhours4days.String
		i.OpenHours.Day4.Hours = openhours4hours.String
		i.OpenHours.Day5.Days = openhours5days.String
		i.OpenHours.Day5.Hours = openhours5hours.String
		i.OpenHours.Day6.Days = openhours6days.String
		i.OpenHours.Day6.Hours = openhours6hours.String
		i.OpenHoursIndividual.Day0.Days = openhoursindividual0days.String
		i.OpenHoursIndividual.Day0.Hours = openhoursindividual0hours.String
		i.OpenHoursIndividual.Day1.Days = openhoursindividual1days.String
		i.OpenHoursIndividual.Day1.Hours = openhoursindividual1hours.String
		i.OpenHoursIndividual.Day2.Days = openhoursindividual2days.String
		i.OpenHoursIndividual.Day2.Hours = openhoursindividual2hours.String
		i.OpenHoursIndividual.Day3.Days = openhoursindividual3days.String
		i.OpenHoursIndividual.Day3.Hours = openhoursindividual3hours.String
		i.OpenHoursIndividual.Day4.Days = openhoursindividual4days.String
		i.OpenHoursIndividual.Day4.Hours = openhoursindividual4hours.String
		i.OpenHoursIndividual.Day5.Days = openhoursindividual5days.String
		i.OpenHoursIndividual.Day5.Hours = openhoursindividual5hours.String
		i.OpenHoursIndividual.Day6.Days = openhoursindividual6days.String
		i.OpenHoursIndividual.Day6.Hours = openhoursindividual6hours.String
		i.SalePointCode = salepointcode.String
		i.Rko = rko.String
		i.Kep = kep.Bool
		i.SuoAvailability = suoavailability.String
		i.HasRamp = hasramp.String
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

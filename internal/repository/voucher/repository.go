package voucherrepository

import (
	voucherdomain "bookcabin-flight-voucher-assignment/internal/domain/voucher"
	"context"
	"fmt"
	"strings"
	"time"
)

func (repository *VoucherRepositoryImpl) GetVouchers(c context.Context, data voucherdomain.Voucher) ([]voucherdomain.Voucher, error) {
	sql := `SELECT id, crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at
		FROM vouchers`

	var (
		conds []string
		vals  []any
		datas []voucherdomain.Voucher
	)

	if len(data.FlightNumber) > 0 {
		conds = append(conds, fmt.Sprintf("flight_number = $%d", len(vals)+1))
		vals = append(vals, data.FlightNumber)
	}

	if len(data.FlightDate) > 0 {
		conds = append(conds, fmt.Sprintf("flight_date = $%d", len(vals)+1))
		vals = append(vals, data.FlightDate)
	}

	if len(data.AircraftType) > 0 {
		conds = append(conds, fmt.Sprintf("aircraft_type = $%d", len(vals)+1))
		vals = append(vals, data.AircraftType)
	}

	if len(conds) > 0 {
		sql += " WHERE " + strings.Join(conds, " AND ")
	}

	rows, err := repository.SQLite.QueryContext(c, sql, vals...)
	if err != nil {
		return datas, err
	}

	defer rows.Close()

	for rows.Next() {
		var v voucherdomain.Voucher
		err := rows.Scan(
			&v.ID,
			&v.CrewName,
			&v.CrewID,
			&v.FlightNumber,
			&v.FlightDate,
			&v.AircraftType,
			&v.Seat1,
			&v.Seat2,
			&v.Seat3,
			&v.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		datas = append(datas, v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return datas, nil
}

func (repository *VoucherRepositoryImpl) CreateVoucher(c context.Context, data voucherdomain.Voucher) error {
	query := `
		INSERT INTO vouchers (
			crew_name,
			crew_id,
			flight_number,
			flight_date,
			aircraft_type,
			seat1,
			seat2,
			seat3,
			created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := repository.SQLite.ExecContext(
		c,
		query,
		data.CrewName,
		data.CrewID,
		data.FlightNumber,
		data.FlightDate,
		data.AircraftType,
		data.Seat1,
		data.Seat2,
		data.Seat3,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return err
	}

	return nil
}

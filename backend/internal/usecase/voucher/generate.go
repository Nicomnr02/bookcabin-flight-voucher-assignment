package voucherusecase

import (
	voucherdomain "bookcabin-flight-voucher-assignment/internal/domain/voucher"
	voucherdto "bookcabin-flight-voucher-assignment/internal/dto/voucher"
	"bookcabin-flight-voucher-assignment/internal/exception"
	"bookcabin-flight-voucher-assignment/pkg/format"
	"bookcabin-flight-voucher-assignment/pkg/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func (usecase *VoucherUsecaseImpl) Generate(c context.Context, request voucherdto.GenerateRequest) (voucherdto.GenerateResponse, error) {
	_, err := time.Parse(format.DATEONLY, request.Date)
	if err != nil {
		return voucherdto.GenerateResponse{}, exception.ErrBadRequest("The date is invalid")
	}

	data := voucherdomain.Voucher{
		FlightNumber: request.FlightNumber,
		FlightDate:   request.Date,
	}

	vouchers, err := usecase.VoucherRP.GetVouchers(c, data)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log.Error().Err(err).Msg(err.Error())
		return voucherdto.GenerateResponse{}, exception.ErrInternalServer("Failed to generate vouchers.")
	}

	if alreadyGenerated := len(vouchers) > 0; alreadyGenerated {
		return voucherdto.GenerateResponse{}, exception.ErrBadRequest("Vouchers have already been generated for the given flight/date.")
	}

	seats := usecase.GetAllSeats()[request.Aircraft]

	randomSeats := usecase.PickRandomSeats(seats, 3)

	voucher := voucherdomain.Voucher{
		CrewName:     request.Name,
		CrewID:       request.ID,
		FlightNumber: request.FlightNumber,
		FlightDate:   request.Date,
		AircraftType: request.Aircraft,
		Seat1:        randomSeats[0],
		Seat2:        randomSeats[1],
		Seat3:        randomSeats[2],
	}

	err = usecase.VoucherRP.CreateVoucher(c, voucher)
	if err != nil {
		logger.Log.Error().Err(err).Msg(err.Error())
		return voucherdto.GenerateResponse{}, exception.ErrInternalServer("Failed to generate vouchers.")
	}

	return voucherdto.GenerateResponse{
		Seats: randomSeats,
	}, nil
}

func (usecase *VoucherUsecaseImpl) GetAllSeats() map[string][]string {
	var seatMaps = map[string][]string{
		"ATR":            usecase.GenerateSeats(1, 18, []string{"A", "C", "D", "F"}),
		"Airbus 320":     usecase.GenerateSeats(1, 32, []string{"A", "B", "C", "D", "E", "F"}),
		"Boeing 737 Max": usecase.GenerateSeats(1, 32, []string{"A", "B", "C", "D", "E", "F"}),
	}
	return seatMaps
}

func (usecase *VoucherUsecaseImpl) GenerateSeats(start, end int, letters []string) []string {
	var seats []string
	for i := start; i <= end; i++ {
		for _, letter := range letters {
			seats = append(seats, fmt.Sprintf("%d%s", i, letter))
		}
	}
	return seats
}

func (usecase *VoucherUsecaseImpl) PickRandomSeats(seats []string, count int) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(seats), func(i, j int) {
		seats[i], seats[j] = seats[j], seats[i]
	})
	return seats[:count]
}

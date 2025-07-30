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
	"time"
)

func (usecase *VoucherUsecaseImpl) Check(c context.Context, request voucherdto.CheckRequest) (voucherdto.CheckResponse, error) {
	_, err := time.Parse(format.DATEONLY, request.Date)
	if err != nil {
		return voucherdto.CheckResponse{}, exception.ErrBadRequest("The date is invalid")
	}

	data := voucherdomain.Voucher{
		FlightNumber: request.FlightNumber,
		FlightDate:   request.Date,
	}

	vouchers, err := usecase.VoucherRP.GetVouchers(c, data)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log.Error().Err(err).Msg(err.Error())
		return voucherdto.CheckResponse{}, exception.ErrInternalServer("Failed to check vouchers.")
	}

	return voucherdto.CheckResponse{Exists: len(vouchers) > 0}, nil
}

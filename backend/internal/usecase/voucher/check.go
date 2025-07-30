package voucherusecase

import (
	voucherdomain "bookcabin-flight-voucher-assignment/internal/domain/voucher"
	voucherdto "bookcabin-flight-voucher-assignment/internal/dto/voucher"
	"bookcabin-flight-voucher-assignment/internal/exception"
	"context"
	"database/sql"
	"errors"
)

func (usecase *VoucherUsecaseImpl) Check(c context.Context, request voucherdto.CheckRequest) (voucherdto.CheckResponse, error) {
	data := voucherdomain.Voucher{
		FlightNumber: request.FlightNumber,
		FlightDate:   request.Date,
	}

	vouchers, err := usecase.VoucherRP.GetVouchers(c, data)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return voucherdto.CheckResponse{}, exception.ErrInternalServer("Failed to check vouchers.")
	}

	return voucherdto.CheckResponse{Exists: len(vouchers) > 0}, nil
}

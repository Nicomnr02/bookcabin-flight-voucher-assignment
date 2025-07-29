package voucherusecase

import (
	voucherdto "bookcabin-flight-voucher-assignment/internal/dto/voucher"
	voucherrepository "bookcabin-flight-voucher-assignment/internal/repository/voucher"
	"context"
)

type VoucherUsecase interface {
	Generate(c context.Context, request voucherdto.GenerateRequest) (voucherdto.GenerateResponse, error)
	Check(c context.Context, request voucherdto.CheckRequest) (voucherdto.CheckResponse, error)
}

type VoucherUsecaseImpl struct {
	VoucherRP voucherrepository.VoucherRepository
}

func New(voucherRP voucherrepository.VoucherRepository) VoucherUsecase {
	return &VoucherUsecaseImpl{voucherRP}
}

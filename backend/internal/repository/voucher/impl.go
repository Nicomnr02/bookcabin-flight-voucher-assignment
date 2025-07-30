package voucherrepository

import (
	voucherdomain "bookcabin-flight-voucher-assignment/internal/domain/voucher"
	"context"
	"database/sql"
)

//go:generate mockery --name VoucherRepository --outpkg vouchermocks --output ../mocks/voucher
type VoucherRepository interface {
	GetVouchers(c context.Context, data voucherdomain.Voucher) ([]voucherdomain.Voucher, error)
	CreateVoucher(c context.Context, data voucherdomain.Voucher) error
}

type VoucherRepositoryImpl struct {
	SQLite *sql.DB
}

func New(SQLite *sql.DB) VoucherRepository {
	return &VoucherRepositoryImpl{
		SQLite: SQLite,
	}
}

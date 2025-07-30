package voucherusecase

import (
	voucherdto "bookcabin-flight-voucher-assignment/internal/dto/voucher"
	vouchermocks "bookcabin-flight-voucher-assignment/internal/repository/mocks/voucher"
	"context"
	"reflect"
	"testing"
)

func TestVoucherUsecaseImpl_Check(t *testing.T) {

	c := context.Background()

	voucherRP := vouchermocks.NewVoucherRepository(t)

	type args struct {
		c       context.Context
		request voucherdto.CheckRequest
	}
	tests := []struct {
		name    string
		usecase *VoucherUsecaseImpl
		args    args
		want    voucherdto.CheckResponse
		wantErr bool
		setMock func()
	}{
		{
			name:    "When date format is invalid throw error",
			usecase: &VoucherUsecaseImpl{voucherRP},
			args: args{
				c: c,
				request: voucherdto.CheckRequest{
					Date: "20-20-2001",
				},
			},
			wantErr: true,
			setMock: func() {
				// voucherRP.On("GetVouchers", voucherdomain.Voucher{FlightDate: "20-20-2001"}).Return(nil, nil)
			},
		},
	}
	for _, tt := range tests {
		tt.setMock()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.usecase.Check(tt.args.c, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("VoucherUsecaseImpl.Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoucherUsecaseImpl.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

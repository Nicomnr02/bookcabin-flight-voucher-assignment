package voucherusecase

import (
	voucherdto "bookcabin-flight-voucher-assignment/internal/dto/voucher"
	vouchermocks "bookcabin-flight-voucher-assignment/internal/repository/mocks/voucher"
	"context"
	"reflect"
	"testing"
)

func TestVoucherUsecaseImpl_Generate(t *testing.T) {

	c := context.Background()

	voucherRP := vouchermocks.NewVoucherRepository(t)

	type args struct {
		c       context.Context
		request voucherdto.GenerateRequest
	}
	tests := []struct {
		name    string
		usecase *VoucherUsecaseImpl
		args    args
		want    voucherdto.GenerateResponse
		wantErr bool
		setMock func()
	}{
		{
			name:    "When date format is invalid throw error",
			usecase: &VoucherUsecaseImpl{voucherRP},
			args: args{
				c: c,
				request: voucherdto.GenerateRequest{
					Date: "20-20-2001",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.usecase.Generate(tt.args.c, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("VoucherUsecaseImpl.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoucherUsecaseImpl.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoucherUsecaseImpl_GetAllSeats(t *testing.T) {
	tests := []struct {
		name    string
		usecase *VoucherUsecaseImpl
		want    map[string][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.usecase.GetAllSeats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoucherUsecaseImpl.GetAllSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoucherUsecaseImpl_GenerateSeats(t *testing.T) {
	type args struct {
		start   int
		end     int
		letters []string
	}
	tests := []struct {
		name    string
		usecase *VoucherUsecaseImpl
		args    args
		want    []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.usecase.GenerateSeats(tt.args.start, tt.args.end, tt.args.letters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoucherUsecaseImpl.GenerateSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoucherUsecaseImpl_PickRandomSeats(t *testing.T) {
	type args struct {
		seats []string
		count int
	}
	tests := []struct {
		name    string
		usecase *VoucherUsecaseImpl
		args    args
		want    []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.usecase.PickRandomSeats(tt.args.seats, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoucherUsecaseImpl.PickRandomSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

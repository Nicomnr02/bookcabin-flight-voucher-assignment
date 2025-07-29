package voucherhandler

import (
	voucherusecase "bookcabin-flight-voucher-assignment/internal/usecase/voucher"

	"github.com/gofiber/fiber/v2"
)

type VoucherHandlerImpl struct {
	fiber     *fiber.App
	VoucherUC voucherusecase.VoucherUsecase
}

func New(fiber *fiber.App, voucherUC voucherusecase.VoucherUsecase) VoucherHandlerImpl {
	return VoucherHandlerImpl{fiber, voucherUC}
}

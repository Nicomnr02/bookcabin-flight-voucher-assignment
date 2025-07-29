package voucherhandler

import (
	voucherdto "bookcabin-flight-voucher-assignment/internal/dto/voucher"
	"bookcabin-flight-voucher-assignment/internal/exception"

	"github.com/gofiber/fiber/v2"
)

func (handler *VoucherHandlerImpl) Router() {
	voucher := handler.fiber.Group("/api")
	voucher.Post("/check", handler.check)
	voucher.Post("/generate", handler.generate)
}

func (handler *VoucherHandlerImpl) check(ctx *fiber.Ctx) error {
	var request voucherdto.CheckRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return exception.ErrorHandler(ctx, exception.ErrInternalServer("Failed to check vouchers"))
	}

	response, err := handler.VoucherUC.Check(ctx.Context(), request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return exception.Response(ctx, fiber.StatusOK, response)
}

func (handler *VoucherHandlerImpl) generate(ctx *fiber.Ctx) error {
	var request voucherdto.GenerateRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return exception.ErrorHandler(ctx, exception.ErrInternalServer("Failed to generate vouchers"))
	}

	response, err := handler.VoucherUC.Generate(ctx.Context(), request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return exception.Response(ctx, fiber.StatusCreated, response)
}

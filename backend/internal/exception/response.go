package exception

import "github.com/gofiber/fiber/v2"

type response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(response{
		Code:    code,
		Success: true,
		Message: "success",
		Data:    data,
	})
}

func ResponseError(ctx *fiber.Ctx, status int, code int, message string) error {
	return ctx.Status(status).JSON(response{
		Code:    code,
		Success: false,
		Message: message,
		Data:    nil,
	})
}

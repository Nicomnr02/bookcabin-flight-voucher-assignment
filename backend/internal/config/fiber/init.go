package fiberconfig

import (
	"bookcabin-flight-voucher-assignment/internal/exception"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler:          exception.ErrorHandler,
		DisableStartupMessage: false,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		BodyLimit:             (40 * 1024 * 1024),
	})

	app.Use(cors.New())

	return app
}

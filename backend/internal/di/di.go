package di

import (
	voucherhandler "bookcabin-flight-voucher-assignment/internal/handler/voucher"
	voucherrepository "bookcabin-flight-voucher-assignment/internal/repository/voucher"
	voucherusecase "bookcabin-flight-voucher-assignment/internal/usecase/voucher"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run(fiber *fiber.App, SQLite *sql.DB) {
	voucherRP := voucherrepository.New(SQLite)
	voucherUC := voucherusecase.New(voucherRP)
	voucherHD := voucherhandler.New(fiber, voucherUC)

	voucherHD.Router()

	err := fiber.Listen("0.0.0.0:5000")
	if err != nil {
		log.Fatal("Failed to run server: ", err.Error())
	}
}

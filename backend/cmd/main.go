package main

import (
	fiberconfig "bookcabin-flight-voucher-assignment/internal/config/fiber"
	sqliteconfig "bookcabin-flight-voucher-assignment/internal/config/sqlite"
	"bookcabin-flight-voucher-assignment/internal/di"
	"bookcabin-flight-voucher-assignment/pkg/logger"
)

func main() {
	logger.Init(true)

	fiber := fiberconfig.Init()
	SQLite := sqliteconfig.Init(false)
	di.Run(fiber, SQLite)
}

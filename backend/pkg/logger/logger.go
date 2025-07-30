package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Init(debug bool) {
	zerolog.TimeFieldFormat = time.RFC3339

	output := os.Stdout

	l := zerolog.New(output).
		With().
		Timestamp().
		Caller().
		Logger()

	if debug {

		l = l.Level(zerolog.DebugLevel)
	} else {
		l = l.Level(zerolog.InfoLevel)
	}

	Log = l
}

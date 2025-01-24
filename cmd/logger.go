package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger = SetupLogger()

func SetupLogger() *zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Logger()

	return &logger
}

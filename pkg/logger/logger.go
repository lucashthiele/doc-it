package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var once sync.Once

var logger zerolog.Logger

func Get() *zerolog.Logger {
	once.Do(func() {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Logger()
	})

	return &logger
}

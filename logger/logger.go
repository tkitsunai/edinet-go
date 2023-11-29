package logger

import (
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

var (
	once   = sync.Once{}
	Logger = zerolog.Logger{}
)

func init() {
	once.Do(func() {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	})
}

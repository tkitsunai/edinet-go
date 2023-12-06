package logger

import (
	"github.com/gofiber/fiber/v2"
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
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	})
}

func RequestLogging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		Logger.Info().Str("method", c.Method()).Str("path", c.Path()).Msg("Request received")
		return c.Next()
	}
}

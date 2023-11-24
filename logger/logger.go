package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strings"
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
			FormatLevel: func(i interface{}) string {
				return strings.ToUpper(fmt.Sprintf("[%s]", i))
			},
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("%s", i)
			},
			FormatCaller: func(i interface{}) string {
				return filepath.Base(fmt.Sprintf("%s", i))
			},
			PartsExclude: []string{
				zerolog.TimestampFieldName,
			},
		}
		Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	})
}

package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type DefaultLogger struct {
	Logger
	isSentry bool
}

var defaultLogger *DefaultLogger

func SetDefaultLogger() {
	var (
		level  = defaultLogLevel
		maxDir = defaultLogCallerMaxDirectory
	)

	level = parseInt(level, os.Getenv("LOG_LEVEL"))
	maxDir = parseInt(maxDir, os.Getenv("LOG_MAX_DIRECTORY"))

	cw := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatMessage: func(i any) string {
			return fmt.Sprintf("\n%s", i)
		},
		FormatCaller: func(i any) string {
			dir, file := filepath.Split(fmt.Sprintf("%s", i))
			list := strings.Split(dir, "/")

			if len(list) < maxDir {
				return fmt.Sprintf("%s%s", dir, file)
			}

			return fmt.Sprintf("%s%s", strings.Join(list[len(list)-maxDir:], "/"), file)
		},
	}

	log := zerolog.New(cw).
		Level(parseLevelZerolog(level)).
		With().
		Timestamp().
		CallerWithSkipFrameCount(defaultLogCallerSkipFrame).
		Logger()

	defaultLogger = &DefaultLogger{
		log:   log,
		level: level,
	}
}

func SetSentry() {
	defaultLogger.isSentry = true
}

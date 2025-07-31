package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var defaultLogger *Logger

func SetDefaultLogger() {
	var (
		level       = defaultLogLevel
		maxDir      = defaultLogCallerMaxDirectory
		sentryLevel = defaultSentryLevel
	)

	level = parseInt(level, os.Getenv("LOG_LEVEL"))
	maxDir = parseInt(maxDir, os.Getenv("LOG_MAX_DIRECTORY"))
	sentryLevel = parseInt(sentryLevel, os.Getenv("SENTRY_LEVEL"))

	cw := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("\n%s", i)
		},
		FormatCaller: func(i interface{}) string {
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

	defaultLogger = &Logger{
		log:         log,
		level:       level,
		sentryLevel: sentryLevel,
	}
}

func SetSentry() {
	defaultLogger.isSentry = true
}

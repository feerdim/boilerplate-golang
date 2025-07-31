package log

import (
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

const (
	debugLevel = iota
	infoLevel
	warnLevel
	errorLevel
	fatalLevel
)

func parseLevelZerolog(level int) zerolog.Level {
	switch level {
	case debugLevel:
		return zerolog.DebugLevel
	case infoLevel:
		return zerolog.InfoLevel
	case warnLevel:
		return zerolog.WarnLevel
	case errorLevel:
		return zerolog.ErrorLevel
	case fatalLevel:
		return zerolog.FatalLevel
	}

	return zerolog.TraceLevel
}

func parseLevelSentry(level int) sentry.Level {
	switch level {
	case debugLevel:
		return sentry.LevelDebug
	case infoLevel:
		return sentry.LevelInfo
	case warnLevel:
		return sentry.LevelWarning
	case errorLevel:
		return sentry.LevelError
	case fatalLevel:
		return sentry.LevelFatal
	}

	return sentry.LevelDebug
}

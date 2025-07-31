package log

import (
	"context"
	"runtime"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type Logger struct {
	log         zerolog.Logger
	level       int
	isSentry    bool
	sentry      *sentry.Event
	sentryLevel int
}

func WithContext(ctx context.Context) *Logger {
	l := Logger{
		log:         defaultLogger.log,
		level:       defaultLogger.level,
		sentryLevel: defaultLogger.sentryLevel,
	}

	if defaultLogger.isSentry {
		l.sentry = sentry.NewEvent()
	}

	l.log.WithContext(ctx)

	return &l
}

func (l *Logger) Debug(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= debugLevel {
		l.log.Debug().Msg(msg)
	}

	if l.sentry != nil && l.sentryLevel <= debugLevel {
		go l.sendSentry(msg, debugLevel)
	}
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= infoLevel {
		l.log.Info().Msg(msg)
	}

	if l.sentry != nil && l.sentryLevel <= infoLevel {
		go l.sendSentry(msg, infoLevel)
	}
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= warnLevel {
		l.log.Warn().Msg(msg)
	}

	if l.sentry != nil && l.sentryLevel <= warnLevel {
		go l.sendSentry(msg, warnLevel)
	}
}

func (l *Logger) Error(err error, msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= errorLevel {
		l.log.Error().Err(err).Msg(msg)
	}

	if l.sentry != nil && l.sentryLevel <= errorLevel {
		pc, file, line, _ := runtime.Caller(1)
		go l.sendExceptionSentry(err, msg, errorLevel, pc, file, line)
	}
}

func (l *Logger) NewError(err, newErr error, fields ...interface{}) error {
	msg := generateMessage(newErr.Error(), fields)

	if l.level <= errorLevel {
		l.log.Error().Err(err).Msg(msg)
	}

	if l.sentry != nil && l.sentryLevel <= errorLevel {
		pc, file, line, _ := runtime.Caller(1)
		go l.sendExceptionSentry(err, msg, errorLevel, pc, file, line)
	}

	return newErr
}

func (l *Logger) Fatal(err error, msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= fatalLevel {
		l.log.Fatal().Err(err).Msg(msg)
	}

	if l.sentry != nil && l.sentryLevel <= fatalLevel {
		pc, file, line, _ := runtime.Caller(1)
		go l.sendExceptionSentry(err, msg, fatalLevel, pc, file, line)
	}
}

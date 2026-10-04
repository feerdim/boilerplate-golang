package log

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type Logger struct {
	log       zerolog.Logger
	level     int
	sentryHub *sentry.Hub
}

func Ctx(ctx context.Context) *Logger {
	l := Logger{
		log:   defaultLogger.log,
		level: defaultLogger.level,
	}

	if defaultLogger.isSentry {
		l.sentryHub = sentry.GetHubFromContext(ctx)
	}

	return &l
}

func (l *Logger) Debug(msg string, fields ...any) {
	msg = generateMessage(msg, fields)

	if l.level <= debugLevel {
		l.log.Debug().Msg(msg)
		l.sendSentry(msg, debugLevel)
	}
}

func (l *Logger) Info(msg string, fields ...any) {
	msg = generateMessage(msg, fields)

	if l.level <= infoLevel {
		l.log.Info().Msg(msg)
		l.sendSentry(msg, infoLevel)
	}
}

func (l *Logger) Warn(msg string, fields ...any) {
	msg = generateMessage(msg, fields)

	if l.level <= warnLevel {
		l.log.Warn().Msg(msg)
		l.sendSentry(msg, warnLevel)
	}
}

func (l *Logger) Error(err error, msg string, fields ...any) {
	msg = generateMessage(msg, fields)

	if l.level <= errorLevel {
		l.log.Error().Err(err).Msg(msg)

		pc, file, line, _ := runtime.Caller(1)
		go l.sendExceptionSentry(err, msg, errorLevel, pc, file, line)
	}
}

func (l *Logger) NewError(err, newErr error, fields ...any) error {
	msg := generateMessage(newErr.Error(), fields)

	if l.level <= errorLevel {
		l.log.Error().Err(err).Msg(msg)

		pc, file, line, _ := runtime.Caller(1)
		go l.sendExceptionSentry(err, msg, errorLevel, pc, file, line)
	}

	return newErr
}

func (l *Logger) Fatal(err error, msg string, fields ...any) {
	msg = generateMessage(msg, fields)

	if l.level <= fatalLevel {
		l.log.Fatal().Err(err).Msg(msg)

		pc, file, line, _ := runtime.Caller(1)
		go l.sendExceptionSentry(err, msg, fatalLevel, pc, file, line)
	}
}

func (l *Logger) sendSentry(msg string, level int) {
	if l.sentryHub == nil {
		return
	}

	l.sentryHub.AddBreadcrumb(&sentry.Breadcrumb{
		Level:   parseLevelSentry(level),
		Message: msg,
	}, nil)
}

func (l *Logger) sendExceptionSentry(err error, msg string, level int, pc uintptr, file string, line int) {
	if l.sentryHub == nil {
		return
	}

	event := sentry.NewEvent()
	event.Level = parseLevelSentry(level)
	event.Message = msg

	event.SetException(err, 1)

	if len(event.Exception) > 0 && event.Exception[0].Stacktrace != nil && len(event.Exception[0].Stacktrace.Frames) > 0 {
		event.Exception[0].Stacktrace.Frames[0].Function = runtime.FuncForPC(pc).Name()
		event.Exception[0].Stacktrace.Frames[0].Filename = filepath.Base(file)
		event.Exception[0].Stacktrace.Frames[0].AbsPath = file
		event.Exception[0].Stacktrace.Frames[0].Lineno = line
	}

	_ = l.sentryHub.CaptureEvent(event)
}

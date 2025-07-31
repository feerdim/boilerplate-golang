package middleware

import (
	"os"
	"strconv"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/getsentry/sentry-go"
	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

func SentryMiddleware(e *echo.Echo) {
	dsn := os.Getenv("SENTRY_DSN")

	debug, err := strconv.ParseBool(os.Getenv("SENTRY_DEBUG"))
	if err != nil {
		debug = constant.DefaultMdwSentryDebug
	}

	if dsn != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:        dsn,
			Debug:      debug,
			SampleRate: constant.DefaultMdwSentrySampleRate,
		}); err != nil {
			log.PrintError(err, "ERROR init sentry")
			return
		}

		e.Use(sentryEcho.New(sentryEcho.Options{}))
		log.SetSentry()
	}
}

package middleware

import (
	"os"
	"strconv"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v5"
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

		e.Use(sentryecho.New(sentryecho.Options{
			Repanic: true,
		}))

		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c *echo.Context) error {
				if hub := sentryecho.GetHubFromContext(c); hub != nil {
					ctx := sentry.SetHubOnContext(c.Request().Context(), hub)
					c.SetRequest(c.Request().WithContext(ctx))
				}

				return next(c)
			}
		})

		log.SetSentry()
	}
}

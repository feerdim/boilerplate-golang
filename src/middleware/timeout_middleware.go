package middleware

import (
	"os"
	"time"

	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TimeoutMiddleware(e *echo.Echo) {
	timeout, err := time.ParseDuration(os.Getenv("APP_REQUEST_TIMEOUT"))
	if err != nil || timeout <= 0 {
		timeout = constant.DefaultMdwTimeout
	}

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: timeout,
		Skipper: middleware.DefaultSkipper,
	}))
}

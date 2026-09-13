package middleware

import (
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func RateLimiterMiddleware(e *echo.Echo) {
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(constant.DefaultMdwRateLimiter)))
}

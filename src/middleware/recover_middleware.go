package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func RecoverMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
}

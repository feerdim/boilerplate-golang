package middleware

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CorsMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderAccept, echo.HeaderContentType, constant.DefaultMdwHeaderToken},
	}))
}

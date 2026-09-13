package domain

import (
	"net/http"
	"os"

	"github.com/feerdim/boilerplate-golang/src/middleware"
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/labstack/echo/v5"

	authDomain "github.com/feerdim/boilerplate-golang/src/domain/auth/routes"
	storageDomain "github.com/feerdim/boilerplate-golang/src/domain/storage/routes"
)

func Routes(e *echo.Echo, k *toolkit.Toolkit) {
	middleware.TimeoutMiddleware(e)
	middleware.SentryMiddleware(e)
	middleware.RecoverMiddleware(e)
	middleware.RateLimiterMiddleware(e)
	middleware.CorsMiddleware(e)

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": os.Getenv("APP_NAME") + "is Running",
		})
	})

	authDomain.AddRoutes(e, k)
	storageDomain.AddRoutes(e, k)
}

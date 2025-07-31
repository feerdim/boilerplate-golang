package routes

import (
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/labstack/echo/v4"

	authApp "github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/application"
	iamApp "github.com/feerdim/boilerplate-golang/src/domain/auth/feature/iam/application"
	userApp "github.com/feerdim/boilerplate-golang/src/domain/auth/feature/user/application"
)

func AddRoutes(e *echo.Echo, t *toolkit.Toolkit) {
	authGroup := e.Group("/auth")

	authApp.AddRoutes(e, t)
	iamApp.AddRoutes(authGroup, t)
	userApp.AddRoutes(authGroup, t)
}

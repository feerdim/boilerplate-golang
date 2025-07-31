package application

import (
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/user/service"
	"github.com/feerdim/boilerplate-golang/src/middleware"
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Group, t *toolkit.Toolkit) {
	svc := service.NewService(t.GetDB())
	mdw := middleware.NewAuthMiddleware(t.GetDB())

	userRoutes := g.Group("/user", mdw.ValidateToken)
	userRoutes.GET("", readUserListApp(svc))
	userRoutes.GET("/:guid", readUserDetailApp(svc))
	userRoutes.POST("", createUserApp(svc))
	userRoutes.PUT("/:guid", updateUserApp(svc))
	userRoutes.DELETE("/:guid", deleteUserApp(svc))
}

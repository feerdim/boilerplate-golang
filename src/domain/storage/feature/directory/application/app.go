package application

import (
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/directory/service"
	"github.com/feerdim/boilerplate-golang/src/middleware"
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Group, t *toolkit.Toolkit) {
	svc := service.NewService(t.GetDB())
	mdw := middleware.NewAuthMiddleware(t.GetDB())

	directoryRoutes := g.Group("/directory", mdw.ValidateToken)
	directoryRoutes.GET("", readDirectoryListApp(svc))
	directoryRoutes.GET("/:guid", readDirectoryDetailApp(svc))
	directoryRoutes.POST("", createDirectoryApp(svc))
	directoryRoutes.PUT("/:guid", updateDirectoryApp(svc))
	directoryRoutes.DELETE("/:guid", deleteDirectoryApp(svc))
}

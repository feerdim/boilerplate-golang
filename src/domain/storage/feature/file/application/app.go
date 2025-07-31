package application

import (
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/file/service"
	"github.com/feerdim/boilerplate-golang/src/middleware"
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Group, t *toolkit.Toolkit) {
	svc := service.NewService(t.GetDB())
	mdw := middleware.NewAuthMiddleware(t.GetDB())

	g.GET("/file/open/:path", openFileApp(svc))

	fileRoutes := g.Group("/file", mdw.ValidateToken)
	fileRoutes.GET("", readFileListApp(svc))
	fileRoutes.GET("/:guid", readFileDetailApp(svc))
	fileRoutes.POST("", createFileApp(svc))
	fileRoutes.PUT("/:guid", updateFileApp(svc))
	fileRoutes.DELETE("/:guid", deleteFileApp(svc))
}

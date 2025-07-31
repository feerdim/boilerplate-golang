package application

import (
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/iam/service"
	"github.com/feerdim/boilerplate-golang/src/middleware"
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Group, t *toolkit.Toolkit) {
	svc := service.NewService(t.GetDB())
	mdw := middleware.NewAuthMiddleware(t.GetDB())

	iamRoutes := g.Group("/iam", mdw.ValidateToken)

	permissionGroupRoutes := iamRoutes.Group("/permission-group")
	permissionGroupRoutes.GET("", readPermissionGroupListApp(svc))
	permissionGroupRoutes.GET("/:guid", readPermissionGroupDetailApp(svc))
	permissionGroupRoutes.POST("", createPermissionGroupApp(svc))
	permissionGroupRoutes.PUT("/:guid", updatePermissionGroupApp(svc))
	permissionGroupRoutes.DELETE("/:guid", deletePermissionGroupApp(svc))

	permissionRoutes := iamRoutes.Group("/permission")
	permissionRoutes.GET("", readPermissionListApp(svc))
	permissionRoutes.GET("/:guid", readPermissionDetailApp(svc))
	permissionRoutes.POST("", createPermissionApp(svc))
	permissionRoutes.PUT("/:guid", updatePermissionApp(svc))
	permissionRoutes.DELETE("/:guid", deletePermissionApp(svc))

	roleRoutes := iamRoutes.Group("/role")
	roleRoutes.GET("", readRoleListApp(svc))
	roleRoutes.GET("/:guid", readRoleDetailApp(svc))
	roleRoutes.POST("", createRoleApp(svc))
	roleRoutes.PUT("/:guid", updateRoleApp(svc))
	roleRoutes.DELETE("/:guid", deleteRoleApp(svc))
}

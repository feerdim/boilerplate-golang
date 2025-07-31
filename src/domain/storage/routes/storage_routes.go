package routes

import (
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/labstack/echo/v4"

	directoryApp "github.com/feerdim/boilerplate-golang/src/domain/storage/feature/directory/application"
	fileApp "github.com/feerdim/boilerplate-golang/src/domain/storage/feature/file/application"
)

func AddRoutes(e *echo.Echo, t *toolkit.Toolkit) {
	storageGroup := e.Group("/storage")

	directoryApp.AddRoutes(storageGroup, t)
	fileApp.AddRoutes(storageGroup, t)
}

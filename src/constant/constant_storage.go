package constant

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// error.
var (
	ErrRangeHeaderInvalid = echo.NewHTTPError(http.StatusBadRequest, "invalid range header")

	ErrRequestedRangeNotSatisfiable = echo.NewHTTPError(http.StatusRequestedRangeNotSatisfiable, "requested range not satisfiable")

	ErrFailedStreamFile = echo.NewHTTPError(http.StatusInternalServerError, "failed stream file")
)

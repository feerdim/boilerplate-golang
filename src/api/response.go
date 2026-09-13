package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

type ResponsePayload struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func ResponseOKForErrNoRows(c *echo.Context, err error, msg string) error {
	if !errors.Is(err, sql.ErrNoRows) {
		return ResponseError(c, err, msg)
	}

	return c.JSON(http.StatusOK, responseDataPayload{
		Message: msg,
	})
}

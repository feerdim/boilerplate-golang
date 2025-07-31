package domain

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/labstack/echo/v4"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var echoError *echo.HTTPError

		// if *echo.HTTPError, let echoError handles it
		if errors.As(err, &echoError) {
			_ = c.JSON(echoError.Code, echoError)
		} else {
			appDebug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))

			switch {
			case !appDebug:
				_ = c.JSON(constant.ErrUnknownSource.Code, constant.ErrUnknownSource)
			default:
				_ = c.JSON(http.StatusInternalServerError, map[string]string{
					"message": err.Error(),
				})
			}
		}

		mappingErrorLogger(err, c)
	}
}

func mappingErrorLogger(err error, c echo.Context) {
	if c.Response().Status >= http.StatusInternalServerError {
		log.PrintWarn(err.Error(),
			"path", c.Request().URL.Path,
		)
	}
}

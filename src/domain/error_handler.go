package domain

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/labstack/echo/v5"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(c *echo.Context, err error) {
		if echoError, ok := errors.AsType[*echo.HTTPError](err); ok {
			_ = c.JSON(echoError.Code, echoError)
			return
		}

		if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
			if resp.Committed {
				return
			}
		}

		message := err.Error()

		_, temp, found := strings.Cut(err.Error(), "message=")
		if found {
			message = temp
		}

		var sc echo.HTTPStatusCoder
		if errors.As(err, &sc) {
			if code := sc.StatusCode(); code != 0 {
				_ = c.JSON(code, api.ResponseErrorPayload{
					Message: message,
				})

				return
			}
		}

		appDebug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))
		if !appDebug {
			_ = c.JSON(constant.ErrUnknownSource.Code, constant.ErrUnknownSource)
		} else {
			_ = c.JSON(http.StatusInternalServerError, api.ResponseErrorPayload{
				Message: message,
			})
		}

		mappingErrorLogger(err, c)
	}
}

func mappingErrorLogger(err error, c *echo.Context) {
	if c.Request().Response.StatusCode >= http.StatusInternalServerError {
		log.PrintWarn(err.Error(),
			"path", c.Request().URL.Path,
		)
	}
}

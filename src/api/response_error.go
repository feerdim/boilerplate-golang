package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/labstack/echo/v4"
)

type responseErrorPayload struct {
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message"`
}

func ResponseError(c echo.Context, err error, msg string) error {
	if err == nil {
		return c.JSON(http.StatusBadRequest, responseErrorPayload{
			Error:   nil,
			Message: msg,
		})
	}

	var echoError *echo.HTTPError

	// if *echo.HTTPError, let echo middleware handles it
	if errors.As(err, &echoError) {
		return err
	}

	e := formatError(err)
	if e != nil {
		return c.JSON(http.StatusBadRequest, responseErrorPayload{
			Error:   e,
			Message: msg,
		})
	}

	return err
}

func ResponseErrorValidate(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, responseErrorPayload{
		Error:   formatErrorValidate(err),
		Message: constant.ErrMsgValidate,
	})
}

func ResponseRedirectError(c echo.Context, baseURL, msg string) error {
	return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s/login/error?message=%s", baseURL, msg))
}

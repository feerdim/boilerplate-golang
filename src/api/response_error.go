package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/labstack/echo/v5"
)

type ResponseErrorPayload struct {
	Error   any    `json:"error,omitempty"`
	Message string `json:"message"`
}

func ResponseError(c *echo.Context, err error, msg string) error {
	if err == nil {
		return c.JSON(http.StatusBadRequest, ResponseErrorPayload{
			Error:   nil,
			Message: msg,
		})
	}

	// if *echo.HTTPError, let echo middleware handles it
	if echoError, ok := errors.AsType[*echo.HTTPError](err); ok {
		return c.JSON(echoError.Code, echoError)
	}

	e := formatError(err)
	if e != nil {
		return c.JSON(http.StatusBadRequest, ResponseErrorPayload{
			Error:   e,
			Message: msg,
		})
	}

	return err
}

func ResponseErrorValidate(c *echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, ResponseErrorPayload{
		Error:   formatErrorValidate(err),
		Message: constant.ErrMsgValidate,
	})
}

func ResponseRedirectError(c *echo.Context, baseURL, msg string) error {
	return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s/login/error?message=%s", baseURL, msg))
}

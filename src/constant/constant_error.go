package constant

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// error.
var (
	ErrFailedParseRequest = echo.NewHTTPError(http.StatusBadRequest, "failed to parse request")

	ErrHeaderTokenNotFound = echo.NewHTTPError(http.StatusUnauthorized, "header authorization not found")
	ErrHeaderTokenInvalid  = echo.NewHTTPError(http.StatusUnauthorized, "invalid header token")
	ErrTokenInvalid        = echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	ErrTokenMissing        = echo.NewHTTPError(http.StatusUnauthorized, "missing token")
	ErrTokenExpired        = echo.NewHTTPError(http.StatusUnauthorized, "expired token")
	// ErrTokenUnauthorized   = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized token").

	ErrForbiddenPermission = echo.NewHTTPError(http.StatusForbidden, "your permission is not allowed to access this resource")
	ErrForbiddenRole       = echo.NewHTTPError(http.StatusForbidden, "your role is not allowed to access this resource")

	ErrDataNotFound = echo.NewHTTPError(http.StatusNotFound, "data not found")

	ErrUnknownSource = echo.NewHTTPError(http.StatusInternalServerError, "an error occurred, please try again later")
)

// error message.
const (
	ErrMsgValidate      = "There are some errors in your request"
	ErrMsgUnknownSource = "an error occurred, please try again later"
)

// error form field.
var (
	// password.
	ErrPasswordIncorrect = errors.New("password incorrect")

	// email.
	ErrAccountNotFound        = errors.New("account not found")
	ErrAccountNotHavePassword = errors.New("account does not have password")
	ErrEmailAlreadyExists     = errors.New("this email has been used")
	ErrEmailSuspended         = errors.New("this email has been suspended")

	// code.
	ErrCodeAlreadyExists = errors.New("this code has been used")
)

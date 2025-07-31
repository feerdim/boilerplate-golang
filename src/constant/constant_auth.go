package constant

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// default.
const (
	DefaultTokenLength                  = 60
	DefaultRoleGUIDUserRegistered       = "019288cc-3c8c-7484-ac0d-3362a88ae018"
	DefaultForgotPasswordExpires        = "1h"
	DefaultUserVerificationExpires      = "1h"
	DefaultAccessTokenExpires           = "1h"
	DefaultRefreshTokenExpires          = "24h"
	DefaultForgotPasswordTemplateHTML   = "./src/template/forgot_password.html"
	DefaultUserVerificationTemplateHTML = "./src/template/user_verification.html"
	DefaultForgotPasswordPathURL        = "/reset-password"
	DefaultUserVerificationPathURL      = "/verification"
)

// error.
var (
	ErrForgotPasswordLinkInvalid = echo.NewHTTPError(http.StatusBadRequest, "invalid forgot password link")
	ErrForgotPasswordLinkExpired = echo.NewHTTPError(http.StatusBadRequest, "forgot password link expired")
	ErrUnknownSSOProvider        = echo.NewHTTPError(http.StatusBadRequest, "unknown sso provider")
	ErrVerificationLinkInvalid   = echo.NewHTTPError(http.StatusBadRequest, "invalid verification link")
	ErrVerificationLinkExpired   = echo.NewHTTPError(http.StatusBadRequest, "verification link expired")
)

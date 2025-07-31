package application

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/service"
	"github.com/labstack/echo/v4"
)

func validateForgotPasswordTokenApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ValidateForgotPasswordTokenRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation forgot password token validation request")
			return api.ResponseErrorValidate(c, err)
		}

		_, err = s.ValidateForgotPasswordTokenService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedValidateForgotPasswordToken)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessValidateForgotPasswordToken,
		})
	}
}

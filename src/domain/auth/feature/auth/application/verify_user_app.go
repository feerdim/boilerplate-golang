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

func verifyUserApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.VerifyUserRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation verification request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.VerifyUserService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedVerifyUser)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessVerifyUser,
		})
	}
}

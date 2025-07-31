package application

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/service"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/labstack/echo/v4"
)

func updateProfileApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.UpdateProfileRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation update profile request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.UpdateProfileService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedUpdateProfile)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessUpdateProfile,
		})
	}
}

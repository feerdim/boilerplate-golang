package application

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/service"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/labstack/echo/v4"
)

func readProfileApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		data, err := s.ReadProfileService(c.Request().Context(), auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetProfile)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToProfileResponse(data),
			Message: msgSuccessGetProfile,
		})
	}
}

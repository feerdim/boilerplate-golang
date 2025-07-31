package application

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/service"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/labstack/echo/v4"
)

func logoutApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		err = s.LogoutService(c.Request().Context(), auth.GetAuth(c).GetClaims())
		if err != nil {
			return api.ResponseError(c, err, msgFailedLogout)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessLogout,
		})
	}
}

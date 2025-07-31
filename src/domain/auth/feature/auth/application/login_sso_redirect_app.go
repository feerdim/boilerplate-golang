package application

import (
	"net/http"
	"os"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/service"
	"github.com/labstack/echo/v4"
)

func loginSSORedirectApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var (
			request payload.LoginSSORedirectRequest
			baseURL = os.Getenv("FRONTEND_URL")
		)

		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return api.ResponseRedirectError(c, baseURL, err.Error())
		}

		err = request.DecodeStateData()
		if err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error decode data", "request", request)
			return api.ResponseRedirectError(c, baseURL, err.Error())
		}

		user, err := s.ExchangeSSOCodeToSSOUserService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseRedirectError(c, baseURL, err.Error())
		}

		data, err := s.SyncSSOUserService(c.Request().Context(), user)
		if err != nil {
			return api.ResponseRedirectError(c, baseURL, err.Error())
		}

		return c.Redirect(http.StatusPermanentRedirect, payload.ToLoginSSORedirectResponse(baseURL, data))
	}
}

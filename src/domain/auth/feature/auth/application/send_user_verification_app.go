package application

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/service"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/labstack/echo/v4"
)

func sendUserVerificationApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()

		user, err := auth.GetAuth(c).GetUser()
		if err != nil {
			log.WithContext(ctx).Error(err, "error find user")
			return api.ResponseError(c, err, msgFailedSendUserVerification)
		}

		if user.VerifiedAt.Valid {
			return api.ResponseError(c, err, msgFailedUserAlreadyVerified)
		}

		err = s.SendUserVerificationService(ctx, user.Name, user.Email)
		if err != nil {
			log.WithContext(ctx).Error(err, "error send user verification")
			return api.ResponseError(c, err, msgFailedSendUserVerification)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessSendUserVerification,
		})
	}
}

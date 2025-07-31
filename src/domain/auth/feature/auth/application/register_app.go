package application

import (
	"context"
	"net/http"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/service"
	"github.com/labstack/echo/v4"
)

func registerApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.RegisterRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation register request")
			return api.ResponseErrorValidate(c, err)
		}

		request.IPAddress = c.RealIP()

		session, user, err := s.RegisterService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedRegister)
		}

		go func() {
			ctx := context.Background()

			if err := s.SendUserVerificationService(ctx, request.Name, request.Email); err != nil {
				log.WithContext(ctx).Error(err, msgFailedSendUserVerification)
			}
		}()

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToSessionResponse(session, user),
			Message: msgSuccessRegister,
		})
	}
}

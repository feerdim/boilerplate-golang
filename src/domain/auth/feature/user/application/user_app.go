package application

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/user/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/user/service"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/labstack/echo/v4"
)

func readUserListApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ReadUserListRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation get user list request")
			return api.ResponseErrorValidate(c, err)
		}

		request.Init()

		data, totalData, err := s.ReadUserListService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetUserList)
		}

		return api.ResponsePaginate(c, request.PaginationPayload, totalData, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToUserResponses(data),
			Message: msgSuccessGetUserList,
		})
	}
}

func readUserDetailApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation get user detail request")
			return api.ResponseErrorValidate(c, err)
		}

		data, err := s.ReadUserDetailService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetUserDetail)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToUserResponse(data),
			Message: msgSuccessGetUserDetail,
		})
	}
}

func createUserApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.UserRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation create user request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.CreateUserService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedCreateUser)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusCreated,
			Data:    nil,
			Message: msgSuccessCreateUser,
		})
	}
}

func updateUserApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.UpdateUserRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation update user request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.UpdateUserService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedUpdateUser)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessUpdateUser,
		})
	}
}

func deleteUserApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation delete user request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.DeleteUserService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedDeleteUser)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteUser,
		})
	}
}

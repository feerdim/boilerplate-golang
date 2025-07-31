package application

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/iam/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/iam/service"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/labstack/echo/v4"
)

func readRoleListApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.PaginationPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation get role list request")
			return api.ResponseErrorValidate(c, err)
		}

		request.Init()

		data, totalData, err := s.ReadRoleListService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetRoleList)
		}

		return api.ResponsePaginate(c, request, totalData, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToRoleResponses(data),
			Message: msgSuccessGetRoleList,
		})
	}
}

func readRoleDetailApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation get role detail request")
			return api.ResponseErrorValidate(c, err)
		}

		data, err := s.ReadRoleDetailService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetRoleDetail)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToRoleResponse(data),
			Message: msgSuccessGetRoleDetail,
		})
	}
}

func createRoleApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.RoleRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation create role request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.CreateRoleService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedCreateRole)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusCreated,
			Data:    nil,
			Message: msgSuccessCreateRole,
		})
	}
}

func updateRoleApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.UpdateRoleRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation update role request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.UpdateRoleService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedUpdateRole)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessUpdateRole,
		})
	}
}

func deleteRoleApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation delete role request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.DeleteRoleService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedDeleteRole)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteRole,
		})
	}
}

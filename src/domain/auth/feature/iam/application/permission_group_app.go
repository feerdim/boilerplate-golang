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

func readPermissionGroupListApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.PaginationPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation get permission list request")
			return api.ResponseErrorValidate(c, err)
		}

		request.Init()

		data, err := s.ReadPermissionGroupListService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetPermissionGroupList)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToPermissionGroupResponses(data),
			Message: msgSuccessGetPermissionGroupList,
		})
	}
}

func readPermissionGroupDetailApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation get permission detail request")
			return api.ResponseErrorValidate(c, err)
		}

		data, err := s.ReadPermissionGroupDetailService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetPermissionGroupDetail)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToPermissionGroupResponse(data),
			Message: msgSuccessGetPermissionGroupDetail,
		})
	}
}

func createPermissionGroupApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.PermissionGroupRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation create permission request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.CreatePermissionGroupService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedCreatePermissionGroup)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessCreatePermissionGroup,
		})
	}
}

func updatePermissionGroupApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.UpdatePermissionGroupRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation update permission request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.UpdatePermissionGroupService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedUpdatePermissionGroup)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessUpdatePermissionGroup,
		})
	}
}

func deletePermissionGroupApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation delete permission request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.DeletePermissionGroupService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedDeletePermissionGroup)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeletePermissionGroup,
		})
	}
}

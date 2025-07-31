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

func readPermissionListApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ReadPermissionListRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation get permission list request")
			return api.ResponseErrorValidate(c, err)
		}

		request.Init()

		data, err := s.ReadPermissionListService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetPermissionList)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToPermissionResponses(data),
			Message: msgSuccessGetPermissionList,
		})
	}
}

func readPermissionDetailApp(s *service.Service) echo.HandlerFunc {
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

		data, err := s.ReadPermissionDetailService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetPermissionDetail)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToPermissionResponse(data),
			Message: msgSuccessGetPermissionDetail,
		})
	}
}

func createPermissionApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.PermissionRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation create permission request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.CreatePermissionService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedCreatePermission)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessCreatePermission,
		})
	}
}

func updatePermissionApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.UpdatePermissionRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.WithContext(c.Request().Context()).Error(err, "error validation update permission request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.UpdatePermissionService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedUpdatePermission)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessUpdatePermission,
		})
	}
}

func deletePermissionApp(s *service.Service) echo.HandlerFunc {
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

		err = s.DeletePermissionService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedDeletePermission)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeletePermission,
		})
	}
}

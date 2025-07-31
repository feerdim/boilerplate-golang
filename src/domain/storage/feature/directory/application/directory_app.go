package application

import (
	"net/http"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/directory/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/directory/service"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/labstack/echo/v4"
)

func readDirectoryListApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ReadDirectoryListRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation get directory list request")
			return api.ResponseErrorValidate(c, err)
		}

		request.Init()

		data, totalData, err := s.ReadDirectoryListService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetDirectoryList)
		}

		return api.ResponsePaginate(c, request.PaginationPayload, totalData, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToDirectoryResponses(data),
			Message: msgSuccessGetDirectoryList,
		})
	}
}

func readDirectoryDetailApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation get directory detail request")
			return api.ResponseErrorValidate(c, err)
		}

		data, err := s.ReadDirectoryDetailService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetDirectoryDetail)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToDirectoryResponse(data),
			Message: msgSuccessGetDirectoryDetail,
		})
	}
}

func createDirectoryApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.DirectoryRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation create directory request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.CreateDirectoryService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedCreateDirectory)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusCreated,
			Data:    nil,
			Message: msgSuccessCreateDirectory,
		})
	}
}

func updateDirectoryApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.UpdateDirectoryRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation update directory request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.UpdateDirectoryService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedUpdateDirectory)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessUpdateDirectory,
		})
	}
}

func deleteDirectoryApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation delete directory request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.DeleteDirectoryService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedDeleteDirectory)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteDirectory,
		})
	}
}

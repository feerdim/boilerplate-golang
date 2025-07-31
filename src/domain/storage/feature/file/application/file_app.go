package application

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/file/helper"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/file/payload"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/file/service"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/feerdim/boilerplate-golang/src/toolkit/storage"
	"github.com/feerdim/boilerplate-golang/src/util"
	"github.com/labstack/echo/v4"
)

func readFileListApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ReadFileListRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation get file list request")
			return api.ResponseErrorValidate(c, err)
		}

		request.Init()

		data, totalData, err := s.ReadFileListService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetFileList)
		}

		return api.ResponsePaginate(c, request.PaginationPayload, totalData, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToFileResponses(data),
			Message: msgSuccessGetFileList,
		})
	}
}

func readFileDetailApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation get file detail request")
			return api.ResponseErrorValidate(c, err)
		}

		data, err := s.ReadFileDetailService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedGetFileDetail)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToFileResponse(data),
			Message: msgSuccessGetFileDetail,
		})
	}
}

func createFileApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.FileRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		request.File, err = c.FormFile("file")
		if err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation create file request")
			return api.ResponseErrorValidate(c, err)
		}

		data, err := s.CreateFileService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedCreateFile)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusCreated,
			Data:    payload.ToFileResponse(data),
			Message: msgSuccessCreateFile,
		})
	}
}

func updateFileApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.UpdateFileRequest
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		request.File, err = c.FormFile("file")
		if err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation update file request")
			return api.ResponseErrorValidate(c, err)
		}

		data, err := s.UpdateFileService(c.Request().Context(), request, auth.GetAuth(c).GetClaims().UserGUID)
		if err != nil {
			return api.ResponseError(c, err, msgFailedUpdateFile)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToFileResponse(data),
			Message: msgSuccessUpdateFile,
		})
	}
}

func deleteFileApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request api.GUIDPayload
		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			log.PrintError(err, "error validation delete file request")
			return api.ResponseErrorValidate(c, err)
		}

		err = s.DeleteFileService(c.Request().Context(), request)
		if err != nil {
			return api.ResponseError(c, err, msgFailedDeleteFile)
		}

		return api.ResponseData(c, api.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteFile,
		})
	}
}

func openFileApp(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var (
			request payload.OpenFileRequest
			ctx     = c.Request().Context()
		)

		if err = c.Bind(&request); err != nil {
			err = log.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		data, err := s.OpenFileService(ctx, request)
		if err != nil {
			return
		}

		stg, err := storage.NewStorage(ctx)
		if err != nil {
			log.WithContext(ctx).Error(err, "error init google cloud storage")
			return
		}
		defer util.CloseBuffer(stg.Client)

		if mimeTypes := strings.Split(data.MimeType, "/"); mimeTypes[0] != "video" {
			rc, err := stg.GetFile(ctx, request.Path)
			if err != nil {
				log.WithContext(ctx).Error(err, "error read file from google cloud storage")
				return err
			}

			return c.Stream(http.StatusOK, data.MimeType, rc)
		}

		obj := stg.Client.Bucket(stg.BucketName).Object(request.Path)

		attrs, err := obj.Attrs(ctx)
		if err != nil {
			log.WithContext(ctx).Error(err, "error get object attributes")
			return
		}

		rangeHeader := c.Request().Header.Get("Range")
		if rangeHeader == "" {
			rc, err := obj.NewReader(ctx)
			if err != nil {
				log.WithContext(ctx).Error(err, "error create reader")
				return err
			}
			defer util.CloseBuffer(rc)

			c.Response().Header().Set("Content-Type", data.MimeType)
			c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", attrs.Size))

			return c.Stream(http.StatusOK, data.MimeType, rc)
		}

		start, end, err := helper.GetRangeHeaderHelper(ctx, rangeHeader, attrs)
		if err != nil {
			return
		}

		rc, err := obj.NewRangeReader(ctx, start, end-start+1)
		if err != nil {
			log.WithContext(ctx).Error(err, "error create range reader")
			return constant.ErrFailedStreamFile
		}
		defer util.CloseBuffer(rc)

		c.Response().Header().Set("Content-Type", data.MimeType)
		c.Response().Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, attrs.Size))
		c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", end-start+1))
		c.Response().WriteHeader(http.StatusPartialContent)

		return c.Stream(http.StatusPartialContent, data.MimeType, rc)
	}
}

package service

import (
	"context"
	"database/sql"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/file/helper"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/file/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
	"gorm.io/gorm"
)

func (s *Service) ReadFileListService(
	ctx context.Context,
	request payload.ReadFileListRequest,
) (data []model.File, totalData int64, err error) {
	statement := s.db.Model(&model.File{})

	if request.DirectoryGUID != "" {
		statement = statement.Where("directory_guid = ?", request.DirectoryGUID)
	}

	if request.SetSearch {
		statement = statement.Where("name ILIKE ?", request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Count(&totalData).Error; err != nil {
		log.WithContext(ctx).Error(err, "error count file", "request", request)
		return
	}

	if request.SetPaginate {
		statement = statement.Limit(request.Limit).Offset(request.Offset)
	}

	if err = statement.Find(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find file", "request", request)
		return
	}

	return
}

func (s *Service) ReadFileDetailService(
	ctx context.Context,
	request api.GUIDPayload,
) (data model.File, err error) {
	data = model.File{GUID: request.GUID}

	if err = s.db.First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find file by guid : "+request.GUID)
		return
	}

	return
}

func (s *Service) CreateFileService(
	ctx context.Context,
	request payload.FileRequest,
	userGUID string,
) (data model.File, err error) {
	path, err := helper.UploadFileHelper(ctx, request)
	if err != nil {
		log.WithContext(ctx).Error(err, "error upload file")
		return
	}

	data = model.File{
		GUID:          util.GenerateUUID(),
		DirectoryGUID: request.DirectoryGUID,
		Name:          request.Name,
		Path:          path,
		Size:          request.Size,
		Extension:     request.Extension,
		MimeType:      request.MimeType,
		CreatedBy:     sql.NullString{String: userGUID, Valid: true},
	}

	if request.Description != nil {
		data.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	if err = s.db.Create(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error create file", "file model", data)
		return
	}

	return
}

func (s *Service) UpdateFileService(
	ctx context.Context,
	request payload.UpdateFileRequest,
	userGUID string,
) (data model.File, err error) {
	var oldPath string

	data.GUID = request.GUID

	if err = s.db.First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find file by guid : "+request.GUID)
		return
	}

	data.DirectoryGUID = request.DirectoryGUID
	data.Name = request.Name
	data.Size = request.Size
	data.Extension = request.Extension
	data.MimeType = request.MimeType
	data.UpdatedBy = sql.NullString{String: userGUID, Valid: true}

	if request.Description != nil {
		data.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	err = util.Transaction(ctx, s.db, func(db *gorm.DB) (err error) {
		if request.File != nil {
			oldPath = data.Path

			data.Path, err = helper.UploadFileHelper(ctx, request.FileRequest)
			if err != nil {
				log.WithContext(ctx).Error(err, "error upload file")
				return
			}
		}

		if err = db.Save(&data).Error; err != nil {
			log.WithContext(ctx).Error(err, "error update file", "file model", data)
			return
		}

		return
	})

	if request.File != nil {
		go helper.DeleteFileHelper(context.Background(), oldPath)
	}

	return
}

func (s *Service) DeleteFileService(
	ctx context.Context,
	request api.GUIDPayload,
) (err error) {
	file := model.File{GUID: request.GUID}

	if err = s.db.First(&file).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find file by guid : "+request.GUID)
		return
	}

	if err = s.db.Delete(&file).Error; err != nil {
		log.WithContext(ctx).Error(err, "error delete file by guid : "+request.GUID)
		return
	}

	go helper.DeleteFileHelper(context.Background(), file.Path)

	return
}

func (s *Service) OpenFileService(
	ctx context.Context,
	request payload.OpenFileRequest,
) (data model.File, err error) {
	data = model.File{}

	if err = s.db.Model(&data).Where("path = ?", request.Path).First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find file by path : "+request.Path)
		return
	}

	return
}

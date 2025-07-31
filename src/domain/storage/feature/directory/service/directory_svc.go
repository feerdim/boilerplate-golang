package service

import (
	"context"
	"database/sql"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/domain/storage/feature/directory/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
)

func (s *Service) ReadDirectoryListService(
	ctx context.Context,
	request payload.ReadDirectoryListRequest,
) (data []model.Directory, totalData int64, err error) {
	statement := s.db.Model(&model.Directory{})

	if request.SetSearch {
		statement = statement.Where("name ILIKE ? OR email ILIKE ?", request.Search, request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Count(&totalData).Error; err != nil {
		log.WithContext(ctx).Error(err, "error count directory", "request", request)
		return
	}

	if request.SetPaginate {
		statement = statement.Limit(request.Limit).Offset(request.Offset)
	}

	if err = statement.Find(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find directory", "request", request)
		return
	}

	return
}

func (s *Service) ReadDirectoryDetailService(
	ctx context.Context,
	request api.GUIDPayload,
) (data model.Directory, err error) {
	data = model.Directory{GUID: request.GUID}

	if err = s.db.First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find directory by guid : "+request.GUID)
		return
	}

	return
}

func (s *Service) CreateDirectoryService(
	ctx context.Context,
	request payload.DirectoryRequest,
	userGUID string,
) (err error) {
	directory := model.Directory{
		GUID:      util.GenerateUUID(),
		Name:      request.Name,
		CreatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if request.DirectoryGUID != nil {
		directory.DirectoryGUID = sql.NullString{String: *request.DirectoryGUID, Valid: true}
	}

	if request.Description != nil {
		directory.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	if err = s.db.Create(&directory).Error; err != nil {
		log.WithContext(ctx).Error(err, "error create directory", "directory model", directory)
		return
	}

	return
}

func (s *Service) UpdateDirectoryService(
	ctx context.Context,
	request payload.UpdateDirectoryRequest,
	userGUID string,
) (err error) {
	directory := model.Directory{GUID: request.GUID}

	if err = s.db.First(&directory).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find directory by guid : "+request.GUID)
		return
	}

	directory.Name = request.Name
	directory.UpdatedBy = sql.NullString{String: userGUID, Valid: true}

	if request.DirectoryGUID != nil {
		directory.DirectoryGUID = sql.NullString{String: *request.DirectoryGUID, Valid: true}
	}

	if request.Description != nil {
		directory.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	if err = s.db.Save(&directory).Error; err != nil {
		log.WithContext(ctx).Error(err, "error update directory", "directory model", directory)
		return
	}

	return
}

func (s *Service) DeleteDirectoryService(
	ctx context.Context,
	request api.GUIDPayload,
) (err error) {
	if err = s.db.Delete(&model.Directory{GUID: request.GUID}).Error; err != nil {
		log.WithContext(ctx).Error(err, "error delete directory by guid : "+request.GUID)
		return
	}

	return
}

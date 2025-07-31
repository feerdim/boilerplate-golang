package service

import (
	"context"
	"database/sql"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/iam/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
)

func (s *Service) ReadPermissionListService(
	ctx context.Context,
	request payload.ReadPermissionListRequest,
) (data []model.Permission, err error) {
	statement := s.db.Model(&model.Permission{}).Preload("PermissionGroup")

	if request.PermissionGroupGUID != "" {
		statement = statement.Where("permission_group_guid = ?", request.PermissionGroupGUID)
	}

	if request.SetSearch {
		statement = statement.Where("name ILIKE ?", request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Find(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find permission", "request", request)
		return
	}

	return
}

func (s *Service) ReadPermissionDetailService(
	ctx context.Context,
	request api.GUIDPayload,
) (data model.Permission, err error) {
	data = model.Permission{GUID: request.GUID}

	if err = s.db.Preload("PermissionGroup").First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find permission by guid : "+request.GUID)
		return
	}

	return
}

func (s *Service) CreatePermissionService(
	ctx context.Context,
	request payload.PermissionRequest,
	userGUID string,
) (err error) {
	permission := model.Permission{
		GUID:                util.GenerateUUID(),
		PermissionGroupGUID: request.PermissionGroupGUID,
		Name:                request.Name,
		CreatedBy:           sql.NullString{String: userGUID, Valid: true},
	}

	if request.Description != nil {
		permission.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	if err = s.db.Create(&permission).Error; err != nil {
		log.WithContext(ctx).Error(err, "error create permission", "permission model", permission)
		return
	}

	return
}

func (s *Service) UpdatePermissionService(
	ctx context.Context,
	request payload.UpdatePermissionRequest,
	userGUID string,
) (err error) {
	permission := model.Permission{
		GUID:      request.GUID,
		Name:      request.Name,
		UpdatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if request.Description != nil {
		permission.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	if err = s.db.Updates(&permission).Error; err != nil {
		log.WithContext(ctx).Error(err, "error update permission", "permission model", permission)
		return
	}

	return
}

func (s *Service) DeletePermissionService(
	ctx context.Context,
	request api.GUIDPayload,
) (err error) {
	if err = s.db.Delete(&model.Permission{GUID: request.GUID}).Error; err != nil {
		log.WithContext(ctx).Error(err, "error delete permission by guid : "+request.GUID)
		return
	}

	return
}

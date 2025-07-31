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

func (s *Service) ReadPermissionGroupListService(
	ctx context.Context,
	request api.PaginationPayload,
) (data []model.PermissionGroup, err error) {
	statement := s.db.Model(&model.PermissionGroup{}).Preload("Permissions")

	if request.SetSearch {
		statement = statement.Where("name ILIKE ?", request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Find(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find permission group", "request", request)
		return
	}

	return
}

func (s *Service) ReadPermissionGroupDetailService(
	ctx context.Context,
	request api.GUIDPayload,
) (data model.PermissionGroup, err error) {
	data = model.PermissionGroup{GUID: request.GUID}

	if err = s.db.First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find permission group by guid : "+request.GUID)
		return
	}

	return
}

func (s *Service) CreatePermissionGroupService(
	ctx context.Context,
	request payload.PermissionGroupRequest,
	userGUID string,
) (err error) {
	permissionGroup := model.PermissionGroup{
		GUID:        util.GenerateUUID(),
		Name:        request.Name,
		CreatedBy:   sql.NullString{String: userGUID, Valid: true},
		Permissions: []model.Permission{},
	}

	if request.Description != nil {
		permissionGroup.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	if err = s.db.Create(&permissionGroup).Error; err != nil {
		log.WithContext(ctx).Error(err, "error create permission group", "permission group model", permissionGroup)
		return
	}

	return
}

func (s *Service) UpdatePermissionGroupService(
	ctx context.Context,
	request payload.UpdatePermissionGroupRequest,
	userGUID string,
) (err error) {
	permissionGroup := model.PermissionGroup{
		GUID:      request.GUID,
		Name:      request.Name,
		UpdatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if request.Description != nil {
		permissionGroup.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	if err = s.db.Updates(&permissionGroup).Error; err != nil {
		log.WithContext(ctx).Error(err, "error update permission group", "permission group model", permissionGroup)
		return
	}

	return
}

func (s *Service) DeletePermissionGroupService(
	ctx context.Context,
	request api.GUIDPayload,
) (err error) {
	if err = s.db.Delete(&model.PermissionGroup{GUID: request.GUID}).Error; err != nil {
		log.WithContext(ctx).Error(err, "error delete permission group by guid : "+request.GUID)
		return
	}

	return
}

package service

import (
	"context"
	"database/sql"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/iam/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
	"gorm.io/gorm"
)

func (s *Service) ReadRoleListService(
	ctx context.Context,
	request api.PaginationPayload,
) (roles []model.Role, totalData int64, err error) {
	statement := s.db.Model(&model.Role{})

	if request.SetSearch {
		statement = statement.Where("name ILIKE ?", request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Count(&totalData).Error; err != nil {
		log.WithContext(ctx).Error(err, "error count role", "request", request)
		return
	}

	if request.SetPaginate {
		statement = statement.Limit(request.Limit).Offset(request.Offset)
	}

	if err = statement.Find(&roles).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find role", "request", request)
		return
	}

	for i := range roles {
		var (
			permissionGUID   []string
			permissionGroups []model.PermissionGroup
		)

		if err = s.db.Model(&model.PermissionRole{}).Where("role_guid = ?", roles[i].GUID).Pluck("permission_guid", &permissionGUID).Error; err != nil {
			log.WithContext(ctx).Error(err, "error find permission role")
			return
		}

		if err = s.db.Model(&model.PermissionGroup{}).Preload("Permissions", func(db *gorm.DB) *gorm.DB {
			return db.Where("guid IN ?", permissionGUID)
		}).Find(&permissionGroups).Error; err != nil {
			log.WithContext(ctx).Error(err, "error find permission group")
			return
		}

		roles[i].PermissionGroups = permissionGroups
	}

	return
}

func (s *Service) ReadRoleDetailService(
	ctx context.Context,
	request api.GUIDPayload,
) (data model.Role, err error) {
	data = model.Role{GUID: request.GUID}

	if err = s.db.Preload("Permissions").First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find role by guid : "+request.GUID)
		return
	}

	var (
		permissionGUID   []string
		permissionGroups []model.PermissionGroup
	)

	if err = s.db.Model(&model.PermissionRole{}).Where("role_guid = ?", data.GUID).Pluck("permission_guid", &permissionGUID).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find permission role")
		return
	}

	if err = s.db.Model(&model.PermissionGroup{}).Preload("Permissions", func(db *gorm.DB) *gorm.DB {
		return db.Where("guid IN ?", permissionGUID)
	}).Find(&permissionGroups).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find permission group")
		return
	}

	data.PermissionGroups = permissionGroups

	return
}

func (s *Service) CreateRoleService(
	ctx context.Context,
	request payload.RoleRequest,
	userGUID string,
) (err error) {
	role := model.Role{
		GUID:      util.GenerateUUID(),
		Name:      request.Name,
		CreatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if request.Description != nil {
		role.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	permissions := make([]model.Permission, len(request.PermissionGUID))
	for i := range request.PermissionGUID {
		permissions[i].GUID = request.PermissionGUID[i]
	}

	err = util.Transaction(ctx, s.db, func(db *gorm.DB) (err error) {
		if err = db.Create(&role).Error; err != nil {
			log.WithContext(ctx).Error(err, "error create role", "role model", role)
			return
		}

		if len(permissions) > 0 {
			if err = db.Model(&role).Association("Permissions").Replace(&permissions); err != nil {
				log.WithContext(ctx).Error(err, "error association role permission", "permission model", permissions)
				return
			}
		}

		return
	})

	return
}

func (s *Service) UpdateRoleService(
	ctx context.Context,
	request payload.UpdateRoleRequest,
	userGUID string,
) (err error) {
	role := model.Role{
		GUID:      request.GUID,
		Name:      request.Name,
		UpdatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if request.Description != nil {
		role.Description = sql.NullString{String: *request.Description, Valid: true}
	}

	permissions := make([]model.Permission, len(request.PermissionGUID))
	for i := range request.PermissionGUID {
		permissions[i].GUID = request.PermissionGUID[i]
	}

	err = util.Transaction(ctx, s.db, func(db *gorm.DB) (err error) {
		if err = db.Updates(&role).Error; err != nil {
			log.WithContext(ctx).Error(err, "error update role", "role model", role)
			return
		}

		if err = db.Model(&role).Association("Permissions").Replace(&permissions); err != nil {
			log.WithContext(ctx).Error(err, "error association role permission", "permission model", permissions)
			return
		}

		return
	})

	return
}

func (s *Service) DeleteRoleService(
	ctx context.Context,
	request api.GUIDPayload,
) (err error) {
	if err = s.db.Delete(&model.Role{GUID: request.GUID}).Error; err != nil {
		log.WithContext(ctx).Error(err, "error delete role by guid : "+request.GUID)
		return
	}

	return
}

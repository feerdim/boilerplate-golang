package service

import (
	"context"
	"database/sql"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/api"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/user/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
)

func (s *Service) ReadUserListService(
	ctx context.Context,
	request payload.ReadUserListRequest,
) (data []model.User, totalData int64, err error) {
	statement := s.db.Model(&model.User{}).Preload("Roles")

	if request.RoleGUID != "" {
		statement = statement.Where("role_id = ?", request.RoleGUID)
	}

	if request.SetSearch {
		statement = statement.Where("name ILIKE ? OR email ILIKE ?", request.Search, request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Count(&totalData).Error; err != nil {
		log.WithContext(ctx).Error(err, "error count user", "request", request)
		return
	}

	if request.SetPaginate {
		statement = statement.Limit(request.Limit).Offset(request.Offset)
	}

	if err = statement.Find(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find user", "request", request)
		return
	}

	return
}

func (s *Service) ReadUserDetailService(
	ctx context.Context,
	request api.GUIDPayload,
) (data model.User, err error) {
	data = model.User{GUID: request.GUID}

	if err = s.db.Preload("Roles").First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find user by guid : "+request.GUID)
		return
	}

	return
}

func (s *Service) CreateUserService(
	ctx context.Context,
	request payload.UserRequest,
	userGUID string,
) (err error) {
	password, err := util.GenerateHashPassword(request.Password)
	if err != nil {
		log.WithContext(ctx).Error(err, "error generate hash password : "+request.Password)
		return
	}

	user := model.User{
		GUID:      util.GenerateUUID(),
		Name:      request.Name,
		Email:     request.Email,
		Password:  password,
		CreatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if err = s.db.Create(&user).Error; err != nil {
		log.WithContext(ctx).Error(err, "error create user", "user model", user)
		return
	}

	return
}

func (s *Service) UpdateUserService(
	ctx context.Context,
	request payload.UpdateUserRequest,
	userGUID string,
) (err error) {
	var password string

	if request.Password != "" {
		password, err = util.GenerateHashPassword(request.Password)
		if err != nil {
			log.WithContext(ctx).Error(err, "error generate hash password : "+request.Password)
			return
		}
	}

	user := model.User{
		GUID:      request.GUID,
		Name:      request.Name,
		Password:  password,
		UpdatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if err = s.db.Updates(&user).Error; err != nil {
		log.WithContext(ctx).Error(err, "error update user", "user model", user)
		return
	}

	return
}

func (s *Service) DeleteUserService(
	ctx context.Context,
	request api.GUIDPayload,
) (err error) {
	if err = s.db.Delete(&model.User{GUID: request.GUID}).Error; err != nil {
		log.WithContext(ctx).Error(err, "error delete user by guid : "+request.GUID)
		return
	}

	return
}

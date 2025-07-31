package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/helper"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
	"gorm.io/gorm"
)

func (s *Service) RegisterService(
	ctx context.Context,
	request payload.RegisterRequest,
) (session model.Session, user *model.User, err error) {
	password, err := util.GenerateHashPassword(request.Password)
	if err != nil {
		log.WithContext(ctx).Error(err, "error generate hash password : "+request.Password)
		return
	}

	user = &model.User{
		GUID:        util.GenerateUUID(),
		Name:        request.Name,
		Email:       request.Email,
		Password:    password,
		ActivatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Roles:       []model.Role{{GUID: constant.DefaultRoleGUIDUserRegistered}},
	}

	session, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(user.GUID))
	if err != nil {
		log.PrintError(err, "error generate session model", "session payload", request.ToSessionPayload(user.GUID))
		return
	}

	err = util.Transaction(ctx, s.db, func(db *gorm.DB) (err error) {
		if err = db.Create(&user).Error; err != nil {
			log.WithContext(ctx).Error(err, "error create user", "user model", user)
			err = util.ValidateUnique(err, constant.ErrEmailAlreadyExists)

			return
		}

		if err = db.Create(&session).Error; err != nil {
			log.WithContext(ctx).Error(err, "error create session", "session model", session)
			return
		}

		return
	})

	return
}

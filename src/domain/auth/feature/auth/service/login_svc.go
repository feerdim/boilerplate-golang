package service

import (
	"context"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/helper"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
)

func (s *Service) LoginService(
	ctx context.Context,
	request payload.LoginRequest,
) (session model.Session, user *model.User, err error) {
	if err = s.db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		err = log.WithContext(ctx).NewError(err, constant.ErrAccountNotFound)
		return
	}

	if err = util.CompareHashPassword(request.Password, user.Password); err != nil {
		err = log.PrintNewError(err, constant.ErrPasswordIncorrect)
		return
	}

	sessionPayload := request.ToSessionPayload(user.GUID)

	session, err = helper.GenerateSessionModel(ctx, sessionPayload)
	if err != nil {
		log.PrintError(err, "error generate session model", "session payload", sessionPayload)
		return
	}

	if err = s.db.Create(&session).Error; err != nil {
		log.WithContext(ctx).Error(err, "error create session", "session model", session)
		return
	}

	return
}

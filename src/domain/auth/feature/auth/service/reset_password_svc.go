package service

import (
	"context"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
)

func (s *Service) ResetPasswordService(
	ctx context.Context,
	request payload.ResetPasswordRequest,
) (err error) {
	userTokenValidation, err := s.ValidateForgotPasswordTokenService(ctx, request.ToValidateForgotPasswordTokenRequest())
	if err != nil {
		return
	}

	password, err := util.GenerateHashPassword(request.Password)
	if err != nil {
		log.WithContext(ctx).Error(err, "error generate hash password : "+request.Password)
		return
	}

	var user model.User

	if err = s.db.Model(&user).Where("email = ?", request.Email).Updates(model.User{Password: password}).Error; err != nil {
		log.WithContext(ctx).Error(err, "error update user by email : "+request.Email)
		return
	}

	if err = s.db.Delete(&userTokenValidation).Error; err != nil {
		log.WithContext(ctx).Error(err, "error delete user validation token", "user validation token model", userTokenValidation)
		return
	}

	return
}

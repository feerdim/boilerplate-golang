package service

import (
	"context"
	"time"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
)

func (s *Service) ValidateForgotPasswordTokenService(
	ctx context.Context,
	request payload.ValidateForgotPasswordTokenRequest,
) (userTokenValidation model.UserTokenValidation, err error) {
	if err = s.db.Where("type = ? AND email = ?", model.UserTokenValidationTypeForgotPassword, request.Email).First(&userTokenValidation).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find user forgot password token by email : "+request.Email)
		return
	}

	if userTokenValidation.Token != request.Token {
		err = constant.ErrForgotPasswordLinkInvalid
		return
	}

	if userTokenValidation.ExpiresAt.Before(time.Now()) {
		err = constant.ErrForgotPasswordLinkExpired
		return
	}

	return
}

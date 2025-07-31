package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
	"gorm.io/gorm"
)

func (s *Service) VerifyUserService(
	ctx context.Context,
	request payload.VerifyUserRequest,
) (err error) {
	var userTokenValidation model.UserTokenValidation

	if err = s.db.Where("type = ? AND email = ?", model.UserTokenValidationTypeVerification, request.Email).First(&userTokenValidation).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find user verification token by email : "+request.Email)
		return
	}

	if userTokenValidation.Token != request.Token {
		err = constant.ErrVerificationLinkInvalid
		return
	}

	if userTokenValidation.ExpiresAt.Before(time.Now()) {
		err = constant.ErrVerificationLinkExpired
		return
	}

	err = util.Transaction(ctx, s.db, func(db *gorm.DB) (err error) {
		var user model.User

		if err = db.Model(&user).Where("email = ?", request.Email).Updates(model.User{VerifiedAt: sql.NullTime{Time: time.Now(), Valid: true}}).Error; err != nil {
			log.WithContext(ctx).Error(err, "error update user by id : "+user.GUID)
			return
		}

		if err = db.Delete(&userTokenValidation).Error; err != nil {
			log.WithContext(ctx).Error(err, "error delete user verification token by email : "+request.Email)
			return
		}

		return
	})

	return
}

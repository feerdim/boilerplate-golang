package service

import (
	"context"
	"time"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
	"gorm.io/gorm"
)

func (s *Service) SendUserVerificationService(
	ctx context.Context,
	name, email string,
) (err error) {
	expiresDuration, err := time.ParseDuration(constant.DefaultUserVerificationExpires)
	if err != nil {
		log.WithContext(ctx).Error(err, "error parse duration : "+constant.DefaultUserVerificationExpires)
		return
	}

	token, err := util.GenerateRandomString(constant.DefaultTokenLength)
	if err != nil {
		return
	}

	userTokenValidation := model.UserTokenValidation{
		GUID:      util.GenerateUUID(),
		Email:     email,
		Type:      model.UserTokenValidationTypeVerification,
		Token:     token,
		ExpiresAt: time.Now().Add(expiresDuration),
	}

	sendUserVerificationPayload := payload.ToSendUserVerificationPayload(name, email, token)

	textHTML, err := util.ParseTemplateHTML(constant.DefaultUserVerificationTemplateHTML, sendUserVerificationPayload)
	if err != nil {
		log.WithContext(ctx).Error(err, "error parse template HTML", "send user verification payload", sendUserVerificationPayload)
		return
	}

	err = util.Transaction(ctx, s.db, func(db *gorm.DB) (err error) {
		if err = db.Create(&userTokenValidation).Error; err != nil {
			log.WithContext(ctx).Error(err, "error create user validation token", "user validation token model", userTokenValidation)
			return
		}

		_, _, err = s.mail.SendMail(ctx, "User Verification", textHTML, email)
		if err != nil {
			log.WithContext(ctx).Error(err, "error send user verification mail to : "+email)
			return
		}

		return
	})

	return
}

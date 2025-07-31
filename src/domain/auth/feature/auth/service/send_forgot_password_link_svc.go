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
	"gorm.io/gorm/clause"
)

func (s *Service) SendForgotPasswordLinkService(
	ctx context.Context,
	request payload.SendForgotPasswordLinkRequest,
) (err error) {
	user := model.User{Email: request.Email}

	if err = s.db.Where(&user).First(&user).Error; err != nil {
		err = log.WithContext(ctx).NewError(err, constant.ErrAccountNotFound)
		return
	}

	token, err := util.GenerateRandomString(constant.DefaultTokenLength)
	if err != nil {
		log.WithContext(ctx).Error(err, "error generate random string")
		return
	}

	expiresDuration, err := time.ParseDuration(constant.DefaultForgotPasswordExpires)
	if err != nil {
		log.WithContext(ctx).Error(err, "error parse duration : "+constant.DefaultForgotPasswordExpires)
		return
	}

	userTokenValidation := model.UserTokenValidation{
		GUID:      util.GenerateUUID(),
		Email:     request.Email,
		Type:      model.UserTokenValidationTypeForgotPassword,
		Token:     token,
		ExpiresAt: time.Now().Add(expiresDuration),
	}

	sendForgotPasswordMailPayload := payload.ToSendForgotPasswordMailPayload(user, userTokenValidation)

	textHTML, err := util.ParseTemplateHTML(constant.DefaultForgotPasswordTemplateHTML, sendForgotPasswordMailPayload)
	if err != nil {
		log.WithContext(ctx).Error(err, "error parse template HTML", "send forgot password mail payload", sendForgotPasswordMailPayload)
		return
	}

	err = util.Transaction(ctx, s.db, func(db *gorm.DB) (err error) {
		if err = db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "email"}, {Name: "type"}},
			DoUpdates: clause.AssignmentColumns([]string{"token", "expires_at"}),
		}).Create(&userTokenValidation).Error; err != nil {
			log.WithContext(ctx).Error(err, "error create or update user validation token", "user token validation model", userTokenValidation)
			return
		}

		_, _, err = s.mail.SendMail(ctx, "Forgot Password", textHTML, request.Email)
		if err != nil {
			log.WithContext(ctx).Error(err, "error send forgot password mail to : "+request.Email)
			return
		}

		return
	})

	return
}

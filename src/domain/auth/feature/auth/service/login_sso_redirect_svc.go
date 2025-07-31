package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/helper"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
	"gorm.io/gorm"
)

func (s *Service) ExchangeSSOCodeToSSOUserService(
	ctx context.Context,
	request payload.LoginSSORedirectRequest,
) (data payload.SSOUserPayload, err error) {
	config, err := helper.NewOAuth2Config(request.Data.Provider)
	if err != nil {
		log.PrintError(err, "unknown sso provider : "+request.Data.Provider)
		return
	}

	token, err := config.Exchange(ctx, request.Code)
	if err != nil {
		log.WithContext(ctx).Error(err, "error exchange code : "+request.Code)
		return
	}

	httpClient := config.Client(ctx, token)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, helper.GetOAuth2ProviderURL(), nil)
	if err != nil {
		log.WithContext(ctx).Error(err, "error init sso http request")
		return
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.WithContext(ctx).Error(err, "error do sso http request")
		return
	}

	defer func() {
		if err = res.Body.Close(); err != nil {
			log.WithContext(ctx).Error(err, "error close sso http response body")
		}
	}()

	if res.StatusCode != http.StatusOK {
		err = constant.ErrUnknownSource
		log.WithContext(ctx).Error(err, "sso http request failed")

		return
	}

	data, err = helper.GetSSOUser(res.Body)
	if err != nil {
		log.WithContext(ctx).Error(err, "error get sso user")
		return
	}

	return
}

func (s *Service) SyncSSOUserService(
	ctx context.Context,
	request payload.SSOUserPayload,
) (data model.Session, err error) {
	var (
		user           model.User
		isUserNotFound bool
	)

	err = s.db.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.WithContext(ctx).Error(err, "error find user by email : "+request.Email)
			return
		}

		isUserNotFound = true

		user = model.User{
			GUID:        util.GenerateUUID(),
			Name:        request.Name,
			Email:       request.Email,
			VerifiedAt:  sql.NullTime{Time: time.Now(), Valid: true},
			ActivatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			Roles:       []model.Role{{GUID: constant.DefaultRoleGUIDUserRegistered}},
		}
	}

	data, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(user.GUID))
	if err != nil {
		return
	}

	err = util.Transaction(ctx, s.db, func(db *gorm.DB) (err error) {
		if isUserNotFound {
			if err = db.Create(&user).Error; err != nil {
				log.WithContext(ctx).Error(err, "error create user", "user model", user)
				return
			}
		}

		if err = s.db.Create(&data).Error; err != nil {
			log.WithContext(ctx).Error(err, "error create session", "session model", data)
			return
		}

		return
	})

	return
}

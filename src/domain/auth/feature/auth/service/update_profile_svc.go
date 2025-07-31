package service

import (
	"context"
	"database/sql"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/util"
)

func (s *Service) UpdateProfileService(
	ctx context.Context,
	request payload.UpdateProfileRequest,
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
		GUID:      userGUID,
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

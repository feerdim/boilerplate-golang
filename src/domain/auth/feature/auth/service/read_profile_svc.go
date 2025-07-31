package service

import (
	"context"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/model"
)

func (s *Service) ReadProfileService(
	ctx context.Context,
	userGUID string,
) (data model.User, err error) {
	data = model.User{GUID: userGUID}

	if err = s.db.Preload("Roles.Permissions").First(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find user by guid : "+userGUID)
		return
	}

	return
}

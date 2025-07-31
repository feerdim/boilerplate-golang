package service

import (
	"context"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/session/jwt"
)

func (s *Service) LogoutService(
	ctx context.Context,
	claims *jwt.AccessTokenPayload,
) (err error) {
	if err = s.db.Delete(&model.Session{GUID: claims.GUID}).Error; err != nil {
		log.WithContext(ctx).Error(err, "error delete session by guid : "+claims.GUID)
		return
	}

	return
}

package service

import (
	"context"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/helper"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/session/jwt"
)

func (s *Service) RefreshTokenService(
	ctx context.Context,
	request payload.RefreshTokenRequest,
) (data model.Session, err error) {
	refreshTokenClaims, err := jwt.ClaimsRefreshToken(request.RefreshToken)
	if err != nil {
		log.PrintError(err, "error claims refresh token : "+request.RefreshToken)
		return
	}

	session := model.Session{GUID: refreshTokenClaims.GUID}

	if err = s.db.First(&session).Error; err != nil {
		log.WithContext(ctx).Error(err, "error find session by guid : "+refreshTokenClaims.GUID)
		return
	}

	data, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(session))
	if err != nil {
		log.PrintError(err, "error generate session model", "session payload", request.ToSessionPayload(session))
		return
	}

	if err = s.db.Updates(&data).Error; err != nil {
		log.WithContext(ctx).Error(err, "error update session", "session model", data)
		return
	}

	return
}

package helper

import (
	"context"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"github.com/feerdim/boilerplate-golang/src/model"
	"github.com/feerdim/boilerplate-golang/src/session/jwt"
)

func GenerateSessionModel(
	ctx context.Context,
	request payload.SessionPayload,
) (data model.Session, err error) {
	accessToken, err := jwt.GenerateAccessToken(request.ToAccessTokenRequest())
	if err != nil {
		log.WithContext(ctx).Error(err, "error generate access token")
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(request.ToRefreshTokenRequest())
	if err != nil {
		log.WithContext(ctx).Error(err, "error generate refresh token")
		return
	}

	data = model.Session{
		GUID:                  request.SessionGUID,
		UserGUID:              request.UserGUID,
		AccessToken:           accessToken.Token,
		AccessTokenExpiresAt:  accessToken.ExpiresAt,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiresAt: refreshToken.ExpiresAt,
		IPAddress:             request.IPAddress,
		UserAgent:             request.UserAgent,
	}

	return
}

package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenPayload struct {
	GUID string
}

func GenerateRefreshToken(request RefreshTokenPayload) (response TokenPayload, err error) {
	expiresDuration, err := time.ParseDuration(os.Getenv("AUTH_REFRESH_TOKEN_EXPIRES"))
	if err != nil {
		return
	}

	expiresAt := time.Now().Add(expiresDuration)

	claims := &jwt.MapClaims{
		"jti": request.GUID,
		"exp": expiresAt.Unix(),
	}

	token, err := GenerateJWT(claims, os.Getenv("AUTH_REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		return
	}

	response = TokenPayload{
		Token:     token,
		ExpiresAt: expiresAt,
	}

	return
}

func ClaimsRefreshToken(token string) (response RefreshTokenPayload, err error) {
	claims, err := ClaimsJWT(token, os.Getenv("AUTH_REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		return
	}

	response = RefreshTokenPayload{
		GUID: claims["jti"].(string),
	}

	return
}

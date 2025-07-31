package jwt

import (
	"time"

	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type TokenPayload struct {
	Token     string
	ExpiresAt time.Time
}

func GenerateJWT(claims jwt.Claims, secretKey string) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secretKey))

	return
}

func ClaimsJWT(token, secretKey string) (claims jwt.MapClaims, err error) {
	_, err = jwt.ParseWithClaims(token, &claims,
		func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, errors.Wrapf(err, "Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secretKey), nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			err = constant.ErrTokenExpired
		}

		return
	}

	return
}

func ClaimsUnverifiedJWT(token string) (claims jwt.MapClaims, err error) {
	jwtToken, _, err := new(jwt.Parser).ParseUnverified(token, &claims)
	if err != nil || jwtToken == nil {
		err = constant.ErrTokenInvalid
		return
	}

	return
}

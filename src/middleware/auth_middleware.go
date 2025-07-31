package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/session/auth"
	"github.com/feerdim/boilerplate-golang/src/session/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	auth *auth.Auth
}

func NewAuthMiddleware(db *gorm.DB) *AuthMiddleware {
	a := auth.NewAuth(db)

	return &AuthMiddleware{
		auth: a,
	}
}

func (am *AuthMiddleware) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := parseHeaderToken(c.Request().Header)
		if err != nil {
			log.PrintError(err, "error parse header token")
			return err
		}

		accessTokenClaims, err := jwt.ClaimsAccessToken(token)
		if err != nil {
			log.PrintError(err, "error claims access token")
			return err
		}

		am.auth.SetClaims(&accessTokenClaims)

		err = am.auth.ValidateSession()
		if err != nil {
			err = log.WithContext(c.Request().Context()).NewError(err, constant.ErrTokenMissing)
			return err
		}

		c.Set("auth", *am.auth)

		return next(c)
	}
}

func parseHeaderToken(h http.Header) (token string, err error) {
	headerDataToken := h.Get(constant.DefaultMdwHeaderToken)
	if !strings.Contains(headerDataToken, "Bearer") {
		err = constant.ErrHeaderTokenNotFound
		return
	}

	splitToken := strings.Split(headerDataToken, fmt.Sprintf("%s ", constant.DefaultMdwHeaderBearer))
	if len(splitToken) <= 1 {
		err = constant.ErrHeaderTokenInvalid
		return
	}

	token = splitToken[1]

	return
}

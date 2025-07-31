package auth

import (
	"github.com/feerdim/boilerplate-golang/src/session/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Auth struct {
	db     *gorm.DB
	claims *jwt.AccessTokenPayload
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func GetAuth(c echo.Context) *Auth {
	a, _ := c.Get("auth").(Auth)

	return &a
}

func (a *Auth) GetClaims() *jwt.AccessTokenPayload {
	return a.claims
}

func (a *Auth) SetClaims(claims *jwt.AccessTokenPayload) {
	a.claims = claims
}

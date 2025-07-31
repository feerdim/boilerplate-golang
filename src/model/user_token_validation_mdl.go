package model

import (
	"time"

	"github.com/feerdim/boilerplate-golang/src/constant"
)

type UserTokenValidation struct {
	GUID      string                  `json:"guid" gorm:"primaryKey"`
	Email     string                  `json:"email"`
	Type      UserTokenValidationType `json:"type"`
	Token     string                  `json:"token"`
	ExpiresAt time.Time               `json:"expired_at"`
	CreatedAt time.Time               `json:"created_at"`
}

type UserTokenValidationType string

const (
	UserTokenValidationTypeForgotPassword UserTokenValidationType = "forgot_password"
	UserTokenValidationTypeVerification   UserTokenValidationType = "verification"
)

func (m UserTokenValidationType) Validate() error {
	switch m {
	case UserTokenValidationTypeForgotPassword, UserTokenValidationTypeVerification:
		return nil
	}

	return constant.ErrFailedParseRequest
}

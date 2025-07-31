package payload

import (
	"time"

	"github.com/feerdim/boilerplate-golang/src/model"
)

type SessionResponse struct {
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  *UserPayload `json:"user"`
}

func ToSessionResponse(session model.Session, user *model.User) (response SessionResponse) {
	response.AccessToken = session.AccessToken
	response.AccessTokenExpiresAt = session.AccessTokenExpiresAt
	response.RefreshToken = session.RefreshToken
	response.RefreshTokenExpiresAt = session.RefreshTokenExpiresAt

	if user != nil {
		response.User = &UserPayload{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
	}

	return
}

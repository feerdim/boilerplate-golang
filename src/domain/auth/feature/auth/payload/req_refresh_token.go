package payload

import "github.com/feerdim/boilerplate-golang/src/model"

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (request *RefreshTokenRequest) ToSessionPayload(session model.Session) (
	params SessionPayload,
) {
	params = SessionPayload{
		SessionGUID: session.GUID,
		UserGUID:    session.UserGUID,
	}

	return
}

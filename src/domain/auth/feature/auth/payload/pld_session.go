package payload

import "github.com/feerdim/boilerplate-golang/src/session/jwt"

type SessionPayload struct {
	SessionGUID string `json:"session_guid"`
	UserGUID    string `json:"user_guid"`
	IPAddress   string `json:"ip_address"`
	UserAgent   string `json:"user_agent"`
}

func (request *SessionPayload) ToAccessTokenRequest() (
	params jwt.AccessTokenPayload,
) {
	params = jwt.AccessTokenPayload{
		GUID:     request.SessionGUID,
		UserGUID: request.UserGUID,
	}

	return
}

func (request *SessionPayload) ToRefreshTokenRequest() (
	params jwt.RefreshTokenPayload,
) {
	params = jwt.RefreshTokenPayload{
		GUID: request.SessionGUID,
	}

	return
}

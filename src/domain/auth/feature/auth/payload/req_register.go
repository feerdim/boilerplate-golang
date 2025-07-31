package payload

import "github.com/feerdim/boilerplate-golang/src/util"

type RegisterRequest struct {
	Token     string `json:"token"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	UserAgent string `json:"user_agent"`
	IPAddress string
}

func (request *RegisterRequest) ToSessionPayload(userGUID string) (
	params SessionPayload,
) {
	params = SessionPayload{
		SessionGUID: util.GenerateUUID(),
		UserGUID:    userGUID,
		IPAddress:   request.IPAddress,
		UserAgent:   request.UserAgent,
	}

	return
}

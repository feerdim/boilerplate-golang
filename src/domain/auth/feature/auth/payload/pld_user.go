package payload

import "github.com/feerdim/boilerplate-golang/src/util"

type UserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SSOUserPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (request *SSOUserPayload) ToUserPayload() (
	params UserPayload,
) {
	params.Name = request.Name
	params.Email = request.Email

	return
}

func (request *SSOUserPayload) ToSessionPayload(userGUID string) (
	params SessionPayload,
) {
	params = SessionPayload{
		SessionGUID: util.GenerateUUID(),
		UserGUID:    userGUID,
	}

	return
}

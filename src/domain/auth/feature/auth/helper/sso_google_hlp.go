package helper

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
)

type googleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func (request *googleUser) toSSOUser() (params payload.SSOUserPayload) {
	params = payload.SSOUserPayload{
		Email: request.Email,
		Name:  fmt.Sprintf("%s %s", request.GivenName, request.FamilyName),
	}

	return
}

func getGoogleUser(resBody io.ReadCloser) (data payload.SSOUserPayload, err error) {
	var user googleUser
	if err = json.NewDecoder(resBody).Decode(&user); err != nil {
		return
	}

	data = user.toSSOUser()

	return
}

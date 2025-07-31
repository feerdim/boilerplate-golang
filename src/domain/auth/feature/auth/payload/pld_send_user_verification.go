package payload

import (
	"os"

	"github.com/feerdim/boilerplate-golang/src/constant"
)

type SendUserVerificationPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToSendUserVerificationPayload(name, email, token string) (params SendUserVerificationPayload) {
	params = SendUserVerificationPayload{
		Name:  name,
		Email: email,
		Token: os.Getenv("FRONTEND_URL") + constant.DefaultUserVerificationPathURL + "?email=" + email + "&token=" + token,
	}

	return
}

package payload

import (
	"os"
	"time"

	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/model"
)

type SendForgotPasswordMailPayload struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}

func ToSendForgotPasswordMailPayload(user model.User, userTokenValidation model.UserTokenValidation) (
	params SendForgotPasswordMailPayload,
) {
	params = SendForgotPasswordMailPayload{
		Name:      user.Name,
		Email:     user.Email,
		Token:     os.Getenv("FRONTEND_URL") + constant.DefaultForgotPasswordPathURL + "?email=" + user.Email + "&token=" + userTokenValidation.Token,
		CreatedAt: userTokenValidation.CreatedAt.Format(time.DateTime),
	}

	return
}

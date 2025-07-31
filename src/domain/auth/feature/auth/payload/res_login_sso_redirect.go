package payload

import (
	"fmt"
	"time"

	"github.com/feerdim/boilerplate-golang/src/model"
)

func ToLoginSSORedirectResponse(baseURL string, entity model.Session) (response string) {
	response = fmt.Sprintf("%s/login/success?access_token=%s&access_token_expired_at=%s&refresh_token=%s&refresh_token_expired_at=%s", baseURL, entity.AccessToken, entity.AccessTokenExpiresAt.Format(time.RFC3339), entity.RefreshToken, entity.RefreshTokenExpiresAt.Format(time.RFC3339))

	return
}

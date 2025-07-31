package helper

import (
	"io"
	"os"

	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	providerGoogle = "google"
)

func NewOAuth2Config(provider string) (config oauth2.Config, err error) {
	switch provider {
	case providerGoogle:
		config.ClientID = os.Getenv("OAUTH2_GOOGLE_CLIENT_ID")
		config.ClientSecret = os.Getenv("OAUTH2_GOOGLE_CLIENT_SECRET")
		config.Endpoint = google.Endpoint
		config.Scopes = []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		}
	default:
		err = constant.ErrUnknownSSOProvider
	}

	config.RedirectURL = os.Getenv("APP_URL") + "/auth/token/sso/%s/redirect" + provider

	return
}

func GetOAuth2ProviderURL() (url string) {
	url = "https://www.googleapis.com/oauth2/v2/userinfo"

	return
}

func GetSSOUser(resBody io.ReadCloser) (data payload.SSOUserPayload, err error) {
	data, err = getGoogleUser(resBody)

	return
}

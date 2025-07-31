package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/helper"
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/payload"
)

func (s *Service) LoginSSOService(
	ctx context.Context,
	request payload.LoginSSORequest,
) (url string, err error) {
	var buf bytes.Buffer

	encoder := base64.NewEncoder(base64.StdEncoding, &buf)

	err = json.NewEncoder(encoder).Encode(request)
	if err != nil {
		log.WithContext(ctx).Error(err, "error encode request", "request", request)
		return
	}

	config, err := helper.NewOAuth2Config(request.Provider)
	if err != nil {
		log.PrintError(err, "unknown sso provider "+request.Provider)
		return
	}

	url = config.AuthCodeURL(buf.String())

	return
}

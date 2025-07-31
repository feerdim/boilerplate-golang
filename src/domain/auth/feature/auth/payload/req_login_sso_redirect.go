package payload

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"
)

type LoginSSORedirectRequest struct {
	Code  string `query:"code"`
	State string `query:"state"`
	Data  LoginSSORequest
}

func (request *LoginSSORedirectRequest) DecodeStateData() (err error) {
	var buffer bytes.Buffer

	_, err = io.Copy(&buffer, base64.NewDecoder(base64.StdEncoding, strings.NewReader(request.State)))
	if err != nil {
		return
	}

	err = json.Unmarshal(buffer.Bytes(), &request.Data)
	if err.Error() == "unexpected end of JSON input" {
		err = json.Unmarshal([]byte(buffer.String()+"}"), &request.Data)
	}

	if err != nil {
		return
	}

	return
}

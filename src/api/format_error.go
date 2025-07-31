package api

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/feerdim/boilerplate-golang/src/constant"
	"github.com/feerdim/boilerplate-golang/src/util"
	"github.com/go-playground/validator/v10"
)

func formatError(err error) (e map[string]string) {
	switch {
	case errors.Is(err, constant.ErrPasswordIncorrect):
		e = map[string]string{"password": util.CapitalFirstLetter(err.Error())}
	case errors.Is(err, constant.ErrAccountNotFound),
		errors.Is(err, constant.ErrAccountNotHavePassword),
		errors.Is(err, constant.ErrEmailAlreadyExists),
		errors.Is(err, constant.ErrEmailSuspended):
		e = map[string]string{"email": util.CapitalFirstLetter(err.Error())}
	case errors.Is(err, constant.ErrCodeAlreadyExists):
		e = map[string]string{"code": util.CapitalFirstLetter(err.Error())}
	}

	return
}

func formatErrorValidate(err error) (message map[string]interface{}) {
	var ve validator.ValidationErrors

	ok := errors.As(err, &ve)
	if !ok {
		return
	}

	message = make(map[string]interface{})

	for _, e := range ve {
		switch e.Tag() {
		case "datetime":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = fmt.Sprintf("Field %s must be date & time", convertCase(e.Field(), ' '))
		case "email":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = "Input must be valid email address"
		case "max":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = fmt.Sprintf("Field %s must be less than %s", convertCase(e.Field(), ' '), e.Param())
		case "min":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = fmt.Sprintf("Field %s must be more than %s", convertCase(e.Field(), ' '), e.Param())
		case "required":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = fmt.Sprintf("Field %s can not empty!", convertCase(e.Field(), ' '))
		}
	}

	return
}

func convertCase(t string, c rune) string {
	buf := &bytes.Buffer{}

	for i, r := range t {
		if i > 0 && unicode.IsUpper(r) {
			if t[i-1] != 'I' && r != 'D' {
				buf.WriteRune(c)
			}
		}

		buf.WriteRune(r)
	}

	return buf.String()
}

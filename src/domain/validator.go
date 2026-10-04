package domain

import "github.com/go-playground/validator/v10"

type Validator struct {
	Validator *validator.Validate
}

func newValidator() Validator {
	return Validator{Validator: validator.New()}
}

func (v Validator) Validate(i any) (err error) {
	err = v.Validator.Struct(i)

	return
}

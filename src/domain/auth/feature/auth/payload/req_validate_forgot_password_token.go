package payload

type ValidateForgotPasswordTokenRequest struct {
	Email string `query:"email" validate:"required,email"`
	Token string `query:"token" validate:"required"`
}

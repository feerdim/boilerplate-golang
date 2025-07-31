package payload

type SendForgotPasswordLinkRequest struct {
	Email string `json:"email" validate:"required,email"`
}

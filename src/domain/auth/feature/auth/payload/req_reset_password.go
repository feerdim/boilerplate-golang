package payload

type ResetPasswordRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (request *ResetPasswordRequest) ToValidateForgotPasswordTokenRequest() (
	params ValidateForgotPasswordTokenRequest,
) {
	params = ValidateForgotPasswordTokenRequest{
		Email: request.Email,
		Token: request.Token,
	}

	return
}

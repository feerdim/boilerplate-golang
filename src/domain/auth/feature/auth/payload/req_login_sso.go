package payload

type LoginSSORequest struct {
	Provider string `param:"provider" json:"provider" validate:"required"`
}

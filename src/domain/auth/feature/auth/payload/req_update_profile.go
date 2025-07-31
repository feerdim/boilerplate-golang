package payload

type UpdateProfileRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password"`
}

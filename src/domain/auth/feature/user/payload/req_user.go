package payload

import "github.com/feerdim/boilerplate-golang/src/api"

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	RoleGUID string `json:"role_guid" validate:"required"`
}

type ReadUserListRequest struct {
	api.PaginationPayload
	RoleGUID string `query:"role_guid"`
}

type UpdateUserRequest struct {
	api.GUIDPayload
	UserRequest
}

package payload

import "github.com/feerdim/boilerplate-golang/src/api"

type PermissionGroupRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type UpdatePermissionGroupRequest struct {
	api.GUIDPayload
	PermissionGroupRequest
}

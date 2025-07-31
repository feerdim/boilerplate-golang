package payload

import "github.com/feerdim/boilerplate-golang/src/api"

type RoleRequest struct {
	Name           string   `json:"name" validate:"required"`
	Description    *string  `json:"description"`
	PermissionGUID []string `json:"permission_guid" validate:"required"`
}

type UpdateRoleRequest struct {
	api.GUIDPayload
	RoleRequest
}

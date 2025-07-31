package payload

import "github.com/feerdim/boilerplate-golang/src/api"

type PermissionRequest struct {
	PermissionGroupGUID string  `json:"permission_group_guid"`
	Name                string  `json:"name" validate:"required"`
	Description         *string `json:"description"`
}

type ReadPermissionListRequest struct {
	api.PaginationPayload
	PermissionGroupGUID string `query:"permission_group_guid"`
}

type UpdatePermissionRequest struct {
	api.GUIDPayload
	PermissionRequest
}

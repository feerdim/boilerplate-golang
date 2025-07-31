package payload

import "github.com/feerdim/boilerplate-golang/src/model"

type RoleResponse struct {
	GUID             string                    `json:"guid"`
	Name             string                    `json:"name"`
	Description      *string                   `json:"description"`
	PermissionGroups []PermissionGroupResponse `json:"permission_groups"`
}

func ToRoleResponse(entity model.Role) (response RoleResponse) {
	response.GUID = entity.GUID
	response.Name = entity.Name

	if entity.Description.Valid {
		response.Description = &entity.Description.String
	}

	response.PermissionGroups = ToPermissionGroupResponses(entity.PermissionGroups)

	return
}

func ToRoleResponses(entities []model.Role) (response []RoleResponse) {
	response = make([]RoleResponse, len(entities))

	for i := range entities {
		response[i] = ToRoleResponse(entities[i])
	}

	return
}

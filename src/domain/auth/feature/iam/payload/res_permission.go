package payload

import "github.com/feerdim/boilerplate-golang/src/model"

type PermissionGroup struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Permission struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type PermissionGroupResponse struct {
	PermissionGroup
	Permissions []Permission `json:"permissions"`
}

type PermissionResponse struct {
	Permission
	PermissionGroup PermissionGroup `json:"permission_group"`
}

func ToPermission(entity model.Permission) (response Permission) {
	response.GUID = entity.GUID
	response.Name = entity.Name

	if entity.Description.Valid {
		response.Description = &entity.Description.String
	}

	return
}

func ToPermissionResponse(entity model.Permission) (response PermissionResponse) {
	response.Permission = ToPermission(entity)
	response.PermissionGroup = ToPermissionGroup(entity.PermissionGroup)

	return
}

func ToPermissionResponses(entities []model.Permission) (response []PermissionResponse) {
	response = make([]PermissionResponse, len(entities))

	for i := range entities {
		response[i] = ToPermissionResponse(entities[i])
	}

	return
}

func ToPermissionGroup(entity model.PermissionGroup) (response PermissionGroup) {
	response.GUID = entity.GUID
	response.Name = entity.Name

	if entity.Description.Valid {
		response.Description = &entity.Description.String
	}

	return
}

func ToPermissionGroupResponse(entity model.PermissionGroup) (response PermissionGroupResponse) {
	response.PermissionGroup = ToPermissionGroup(entity)

	response.Permissions = make([]Permission, len(entity.Permissions))

	for i := range entity.Permissions {
		response.Permissions[i] = ToPermission(entity.Permissions[i])
	}

	return
}

func ToPermissionGroupResponses(entities []model.PermissionGroup) (response []PermissionGroupResponse) {
	response = make([]PermissionGroupResponse, len(entities))

	for i := range entities {
		response[i] = ToPermissionGroupResponse(entities[i])
	}

	return
}

package payload

import (
	"time"

	"github.com/feerdim/boilerplate-golang/src/model"
)

type ProfileResponse struct {
	GUID        string         `json:"guid"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	VerifiedAt  *string        `json:"verified_at"`
	ActivatedAt *string        `json:"activated_at"`
	ActivatedBy *string        `json:"activated_by"`
	CreatedAt   string         `json:"created_at"`
	CreatedBy   *string        `json:"created_by"`
	UpdatedAt   *string        `json:"updated_at"`
	UpdatedBy   *string        `json:"updated_by"`
	Roles       []RoleResponse `json:"roles"`
}

type RoleResponse struct {
	GUID        string               `json:"guid"`
	Name        string               `json:"name"`
	Description *string              `json:"description"`
	Permissions []PermissionResponse `json:"permissions"`
}

type PermissionResponse struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func ToProfileResponse(entity model.User) (response ProfileResponse) {
	response.GUID = entity.GUID
	response.Name = entity.Name
	response.Email = entity.Email
	response.CreatedAt = entity.CreatedAt.Format(time.RFC3339)

	if entity.VerifiedAt.Valid {
		verifiedAt := entity.VerifiedAt.Time.Format(time.RFC3339)
		response.VerifiedAt = &verifiedAt
	}

	if entity.ActivatedAt.Valid {
		activeAt := entity.ActivatedAt.Time.Format(time.RFC3339)
		response.ActivatedAt = &activeAt
	}

	if entity.ActivatedBy.Valid {
		response.ActivatedBy = &entity.ActivatedBy.String
	}

	if entity.CreatedBy.Valid {
		response.CreatedBy = &entity.CreatedBy.String
	}

	if entity.UpdatedAt.Valid {
		updatedAt := entity.UpdatedAt.Time.Format(time.RFC3339)
		response.UpdatedAt = &updatedAt
	}

	if entity.UpdatedBy.Valid {
		response.UpdatedBy = &entity.UpdatedBy.String
	}

	response.Roles = make([]RoleResponse, len(entity.Roles))

	for i := range entity.Roles {
		response.Roles[i] = ToRoleResponse(entity.Roles[i])
	}

	return
}

func ToRoleResponse(entity model.Role) (response RoleResponse) {
	response.GUID = entity.GUID
	response.Name = entity.Name

	if entity.Description.Valid {
		response.Description = &entity.Description.String
	}

	response.Permissions = make([]PermissionResponse, len(entity.Permissions))

	for i := range entity.Permissions {
		response.Permissions[i] = ToPermissionResponse(entity.Permissions[i])
	}

	return
}

func ToPermissionResponse(entity model.Permission) (response PermissionResponse) {
	response.GUID = entity.GUID
	response.Name = entity.Name

	if entity.Description.Valid {
		response.Description = &entity.Description.String
	}

	return
}

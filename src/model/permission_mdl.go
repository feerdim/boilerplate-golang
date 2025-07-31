package model

import (
	"database/sql"
	"time"
)

type Permission struct {
	GUID                string          `json:"guid" gorm:"primaryKey"`
	PermissionGroupGUID string          `json:"permission_group_guid"`
	Name                string          `json:"name"`
	Description         sql.NullString  `json:"description"`
	CreatedAt           time.Time       `json:"created_at"`
	CreatedBy           sql.NullString  `json:"created_by"`
	UpdatedAt           sql.NullTime    `json:"updated_at"`
	UpdatedBy           sql.NullString  `json:"updated_by"`
	PermissionGroup     PermissionGroup `json:"permission_group" gorm:"foreignKey:PermissionGroupGUID"`
	Roles               []Role          `json:"roles" gorm:"many2many:permission_role"`
}

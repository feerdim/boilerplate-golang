package model

import (
	"database/sql"
	"time"
)

type Role struct {
	GUID             string            `json:"guid" gorm:"primaryKey"`
	Name             string            `json:"name"`
	Description      sql.NullString    `json:"description"`
	CreatedAt        time.Time         `json:"created_at"`
	CreatedBy        sql.NullString    `json:"created_by"`
	UpdatedAt        sql.NullTime      `json:"updated_at"`
	UpdatedBy        sql.NullString    `json:"updated_by"`
	Permissions      []Permission      `json:"permissions" gorm:"many2many:permission_role"`
	PermissionGroups []PermissionGroup `json:"permission_groups" gorm:"many2many:permission_role"`
	Users            []User            `json:"users" gorm:"many2many:role_user"`
}

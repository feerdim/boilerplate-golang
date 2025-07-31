package model

import (
	"database/sql"
	"time"
)

type User struct {
	GUID        string         `json:"guid" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	VerifiedAt  sql.NullTime   `json:"verified_at"`
	ActivatedAt sql.NullTime   `json:"activated_at"`
	ActivatedBy sql.NullString `json:"activated_by"`
	CreatedAt   time.Time      `json:"created_at"`
	CreatedBy   sql.NullString `json:"created_by"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	UpdatedBy   sql.NullString `json:"updated_by"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
	DeletedBy   sql.NullString `json:"deleted_by"`
	Roles       []Role         `json:"roles" gorm:"many2many:role_user"`
}

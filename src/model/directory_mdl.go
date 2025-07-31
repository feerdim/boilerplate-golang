package model

import (
	"database/sql"
	"time"
)

type Directory struct {
	GUID          string         `json:"guid" gorm:"primaryKey"`
	DirectoryGUID sql.NullString `json:"directory_guid"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	CreatedAt     time.Time      `json:"created_at"`
	CreatedBy     sql.NullString `json:"created_by"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	UpdatedBy     sql.NullString `json:"updated_by"`
	DeletedAt     sql.NullTime   `json:"deleted_at"`
	DeletedBy     sql.NullString `json:"deleted_by"`
	Files         []File         `json:"files" gorm:"foreignKey:DirectoryGUID"`
}

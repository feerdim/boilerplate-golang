package model

import (
	"database/sql"
	"time"
)

type File struct {
	GUID          string         `json:"guid" gorm:"primaryKey"`
	DirectoryGUID string         `json:"directory_guid"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	Path          string         `json:"path"`
	Size          int64          `json:"size"`
	Extension     string         `json:"extension"`
	MimeType      string         `json:"mime_type"`
	CreatedAt     time.Time      `json:"created_at"`
	CreatedBy     sql.NullString `json:"created_by"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	UpdatedBy     sql.NullString `json:"updated_by"`
	DeletedAt     sql.NullTime   `json:"deleted_at"`
	DeletedBy     sql.NullString `json:"deleted_by"`
	Directory     Directory      `json:"directory" gorm:"foreignKey:DirectoryGUID"`
}

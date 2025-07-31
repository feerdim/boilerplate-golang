package model

import (
	"database/sql"
	"time"
)

type Session struct {
	GUID                  string       `json:"guid" gorm:"primaryKey"`
	UserGUID              string       `json:"user_guid"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	IPAddress             string       `json:"ip_address"`
	UserAgent             string       `json:"user_agent"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             sql.NullTime `json:"updated_at"`
}

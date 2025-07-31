package service

import (
	"github.com/feerdim/boilerplate-golang/src/toolkit/mail"
	"gorm.io/gorm"
)

type Service struct {
	db   *gorm.DB
	mail *mail.Mail
}

func NewService(db *gorm.DB, mail *mail.Mail) *Service {
	return &Service{
		db:   db,
		mail: mail,
	}
}

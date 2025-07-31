package toolkit

import (
	"github.com/feerdim/boilerplate-golang/src/toolkit/mail"
	"github.com/feerdim/boilerplate-golang/src/toolkit/storage"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Toolkit struct {
	db  *gorm.DB
	dbx *sqlx.DB
	m   *mail.Mail
	stg *storage.Storage
}

func NewToolkit(
	db *gorm.DB,
	dbx *sqlx.DB,
	m *mail.Mail,
	stg *storage.Storage,
) *Toolkit {
	return &Toolkit{
		db:  db,
		dbx: dbx,
		m:   m,
		stg: stg,
	}
}

func (k *Toolkit) GetDB() *gorm.DB {
	return k.db
}

func (k *Toolkit) GetDBX() *sqlx.DB {
	return k.dbx
}

func (k *Toolkit) GetMail() *mail.Mail {
	return k.m
}

func (k *Toolkit) GetStorage() *storage.Storage {
	return k.stg
}

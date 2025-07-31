package auth

import (
	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/model"
)

func (a *Auth) GetUser() (data model.User, err error) {
	data = model.User{GUID: a.claims.UserGUID}

	if err = a.db.First(&data).Error; err != nil {
		log.PrintError(err, "error find user by guid : "+a.claims.UserGUID)
		return
	}

	return
}

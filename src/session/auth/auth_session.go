package auth

import (
	"github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/model"
)

func (a *Auth) ValidateSession() (err error) {
	data := model.Session{GUID: a.claims.GUID}

	if err = a.db.First(&data).Error; err != nil {
		log.PrintError(err, "error find session by guid : "+a.claims.GUID)
		return
	}

	return
}

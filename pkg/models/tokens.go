package models

import (
	"github.com/bawazy/auth/pkg/config"
)

type Authorizations struct {
	Token string `json:"token"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Authorizations{})
}

func (u *Authorizations) CreateToken() error {
	db.NewRecord(u)
	db.Create(&u)
	return db.Error
}

func GetAuthbyToken(token string) []Authorizations {
	var getToken []Authorizations
	db.Where("token=?", token).Find(&getToken)
	return getToken

}

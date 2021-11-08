package models

import (
	"github.com/bawazy/auth/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"" json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users

}

func GetUserbyId(Id int64) (*User, *gorm.DB) {
	var getUser User
	db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db

}

func GetUserbyUsername(username string) []User {
	var getUser []User
	db.Where("username=?", username).Find(&getUser)
	return getUser

}

func GetUserbyEmail(email string) []User {
	var getUser []User
	db.Where("email=?", email).Find((&getUser))
	return getUser
}

func DeleteUser(Id int64) {
	var user User
	db.Where("ID=?", Id).Delete(&user)
}

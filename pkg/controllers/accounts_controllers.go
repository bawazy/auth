package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bawazy/auth/pkg/models"
	"github.com/bawazy/auth/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserDetails struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)

	existinguser := models.GetUserbyUsername(user.Username)

	if len(existinguser) == 0 {
		w.WriteHeader(http.StatusConflict)
	} else if checkPasswordHash(user.Password, existinguser[0].Password) {

		//create a token
		d, _ := HashPassword(existinguser[0].Username)
		usrDetails := &UserDetails{
			ID:       existinguser[0].ID,
			Username: existinguser[0].Username,
			Password: existinguser[0].Password,
			Email:    existinguser[0].Email,
			Token:    d,
		}

		//token := models.CreateToken()
		//save token to token table
		//T, _ := json.Marshal(token)
		// tkndetails, _ := json.Marshal(token)
		usr, _ := json.Marshal(usrDetails)

		// tkndetails = append(tkndetails, usr...)

		// err := token.CreateToken()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		//res = append(res[:len(res)-1], []byte(token)...)

		w.WriteHeader(http.StatusOK)
		w.Write(usr)

	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	utils.ParseBody(r, user)
	password, err := HashPassword(user.Password)
	user.Password = string(password)
	if err != nil {
		log.Fatal(err)
	}
	existinguser := models.GetUserbyUsername(user.Username)

	if len(existinguser) != 0 {
		w.WriteHeader(http.StatusConflict)
	} else {
		t := user.CreateUser()
		res, _ := json.Marshal(t)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allusers := models.GetAllUsers()
	res, _ := json.Marshal(allusers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func HashPassword(password string) (string, error) {
	a, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(a), err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

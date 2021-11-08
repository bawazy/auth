package routes

import (
	"github.com/bawazy/auth/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/login/", controllers.Login).Methods("POST")
	router.HandleFunc("/register/", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/users/", controllers.GetAllUsers).Methods("GET")
}

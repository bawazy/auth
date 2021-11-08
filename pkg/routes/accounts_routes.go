package routes

import (
	"github.com/bawazy/auth/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/login/", controllers.Login).Methods("POST")
	router.HandleFunc("/register/", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/users/", controllers.GetAllUsers).Methods("GET")
}

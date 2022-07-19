package controllers

import (
	"github.com/gorilla/mux"
	"clean-go/src/main/factories/domain/usecases"
	"clean-go/src/application/controller"
	)


	func MakeUserHandlers(r *mux.Router) {

		userService := usecases.MakeUserService()
		r.Handle("/v1/user",
		controller.GetUser(userService),
		).Methods("GET")
	
	}
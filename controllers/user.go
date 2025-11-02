package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}


func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController)GetUserByID(w http.ResponseWriter,r *http.Request){
	fmt.Println("creating user in user controller")
	uc.UserService.GetUserByID()
	w.Write([]byte("User registration endpoint"))
}
package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
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
	fmt.Println("getting user in user controller")
	uc.UserService.GetUserByID()
	w.Write([]byte("User fetching endpoint"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create User called in UserController")
	uc.UserService.CreateUser()
	w.Write([]byte("user creation endpoint done"))
}

func (uc *UserController) Login(w http.ResponseWriter,r *http.Request){
	fmt.Println("login User called in UserController")

	var payload dto.LoginUserRequestDTO
	jsonErr := utils.ReadJsonBody(r,&payload);
	if jsonErr != nil{
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Something went wrong while logging in",jsonErr)
		return
	}

	validationErr := utils.Validator.Struct(payload)
	if validationErr != nil{
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Invalid input data",validationErr)
		return
	}

	jwtToken,err:=uc.UserService.LoginUser(&payload)
	if err!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed to login user",err)
	}	
	
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User logged in successfully",jwtToken)

}


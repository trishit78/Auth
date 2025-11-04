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
func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// create a dto for payload
	var payload dto.GetUserByIdDTO

	//readjson and give error
	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while getting user by id", jsonErr)
		return
	}

	fmt.Println("Payload received:", payload.Id)
	//validate payload
	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
		return
	}

	user, err := uc.UserService.GetUserByID(payload.Id)

	//deal with this err
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to get user", err)
		return
	}

	//success response
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User got successfully", user)

}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var payload dto.CreateUserDTO

	//readjson and give error
	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while reading the json", jsonErr)
		return
	}

	fmt.Println("Payload received:", payload)
	//validate payload
	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
		return
	}

	user, err := uc.UserService.CreateUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	//success response
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User created successfully", user)
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


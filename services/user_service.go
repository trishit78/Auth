package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserByID() error
	CreateUser() error
	LoginUser(payload *dto.LoginUserRequestDTO) (string,error)


}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserByID() error {
	fmt.Println("Creating user in UserService")
	 u.userRepository.GetByID(3)
	return nil
}

func (u *UserServiceImpl) CreateUser() error{
	fmt.Println("Creating user in userservice")
	password:="password"
	hashedPassword ,err:= utils.HashPassword(password)
	if err!=nil{
		return err
	} 
	u.userRepository.Create("trishit3","trishit4256@gmail.com",hashedPassword)
	return nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO)   (string,error){
	fmt.Println("Login user in userservice")
	password:=payload.Password
	email:=payload.Email
	user,err:= u.userRepository.GetByEmail(email)
	if err!=nil{
		fmt.Println("error fetching the user",err)
		return "", err
	}
	if user == nil{
		fmt.Println("No user found with the given email")
		return "",fmt.Errorf("no user found with email")
	}	

	isPasswordValid := utils.CheckPasswordHash(password,user.Password)
	
	if !isPasswordValid{
		fmt.Println("Password does not match")
		return "",nil
	}
	
	jwtpayload:=jwt.MapClaims{
		"email":user.Email,
		"id":user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwtpayload)
	tokenString,err := token.SignedString([]byte(env.GetString("JWT_SECRET","TOKEN")))

	if err!=nil{
		fmt.Println("Error signing token",err)
		return "",err
	}
	fmt.Println("JWT Token",tokenString)	
	
	return tokenString,nil
}


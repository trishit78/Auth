package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
	env "AuthInGo/config/env"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserByID() error
	CreateUser() error
	LoginUser() (string,error)


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

func (u *UserServiceImpl) LoginUser()   (string,error){
	fmt.Println("Login user in userservice")
	password:="password"
	user,err:= u.userRepository.GetByEmail("trishit4256@gmail.com")
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
	
	payload:=jwt.MapClaims{
		"email":user.Email,
		"id":user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	tokenString,err := token.SignedString([]byte(env.GetString("JWT_SECRET","TOKEN")))

	if err!=nil{
		fmt.Println("Error signing token",err)
		return "",err
	}
	fmt.Println("JWT Token",tokenString)	
	
	return tokenString,nil
}


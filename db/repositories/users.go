package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID() (*models.User,error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository{
	return &UserRepositoryImpl{
		db:db,
	}
}


func (u *UserRepositoryImpl) GetByID() (*models.User,error) {
	fmt.Println("Feetching user in UserRepository")

	query:="SELECT id,username,email,password,created_at,updated_at from users where id= ?"

	row :=u.db.QueryRow(query,3)

	user:= &models.User{}

	err	:=row.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)

	if err!=nil{
		if err == sql.ErrNoRows{
			fmt.Println("no user found with the given ID")
			return nil,err
		} else{
			fmt.Println("Error scanning user",err)
			return nil,err
		}
	}
	fmt.Println("User fetched successfully",user)
	return user,nil
}
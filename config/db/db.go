package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB,error) {
	cfg:=mysql.NewConfig()	
	cfg.User = env.GetString("DB_USER","root")
	cfg.Net = env.GetString("DB_NET","tcp")
	cfg.Passwd = env.GetString("DB_PASSWORD","root")
	cfg.Addr= env.GetString("DB_ADDR","127.0.0.1:3306")
	cfg.DBName = env.GetString("DBName","auth_dev")

	fmt.Println("Connecting to database:",cfg.DBName,cfg.FormatDSN())

	db,err :=sql.Open("mysql",cfg.FormatDSN())
	if err != nil{
		fmt.Println("error connecting to database",err)
		return nil,err
	}
	fmt.Println("Trying to connect to database")
	pingErr:=db.Ping()
	if pingErr!=nil{
		fmt.Println("Error pinging to database",pingErr)
		return  nil,err
	}
	fmt.Println("Connected to database successfully",cfg.DBName)

	return db,nil

}
package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
	dbConfig "AuthInGo/config/db"


)


type Config struct{
	Addr string   //port
}

type Application struct{
	Config Config

}

func NewConfig(addr string) Config{
	port:=config.GetString("PORT",":8080")
	return Config{
		Addr:port,
	}
}

func NewApplication(cfg Config)*Application{
	return &Application{
		Config:cfg,
	}
}



func (app *Application) Run() error{
	db,err:=dbConfig.SetupDB()
	
	if err!=nil{
		fmt.Println("Error setting up database",err)
		return err
	}

	ur:=repo.NewUserRepository(db)
	us:=services.NewUserService(ur)
	uc:=controllers.NewUserController(us)
	uRouter:= router.NewUserRouter(uc)
	
	
	server:=&http.Server{
		Addr: app.Config.Addr,
		Handler: router.SetupRouter(uRouter),
		ReadTimeout: 10 * time.Second,  // set read timeout to 10 sec
		WriteTimeout: 10 * time.Second,  // // set write timeout to 10 sec
	}
	fmt.Println("Starting server on ",app.Config.Addr);
	return server.ListenAndServe()
}

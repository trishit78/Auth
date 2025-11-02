package app

import (
	"fmt"
	"net/http"
	"time"
)


type Config struct{
	Addr string   //port
}

type Application struct{
	Config Config
}

func NewConfig(addr string) Config{
	return Config{
		Addr:addr,
	}
}

func NewApplication(cfg Config)*Application{
	return &Application{
		Config:cfg,
	}
}



func (app *Application) Run() error{
	server:=&http.Server{
		Addr: app.Config.Addr,
		Handler: nil,           // setup a chi router
		ReadTimeout: 10 * time.Second,  // set read timeout to 10 sec
		WriteTimeout: 10 * time.Second,  // // set write timeout to 10 sec
	}
	fmt.Println("Starting server on ",app.Config.Addr);
	return server.ListenAndServe()
}

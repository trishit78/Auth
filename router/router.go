package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"github.com/go-chi/chi/v5"
)

type Router interface{
	Register(r chi.Router)
}


func SetupRouter(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Use(middlewares.RequestLogger)

	chiRouter.Get("/ping",controllers.PingController)
	
	UserRouter.Register(chiRouter)
	return chiRouter
}


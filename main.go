package main

import (
	"debugpedia-api/controller"
	"debugpedia-api/db"
	"debugpedia-api/repository"
	"debugpedia-api/router"
	"debugpedia-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	debugRepository := repository.NewDebugRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	debugUsecase := usecase.NewDebugUsecase(debugRepository)
	userController := controller.NewUserController(userUsecase)
	debugController := controller.NewDebugController(debugUsecase)
	e := router.NewRouter(userController, debugController)
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"debugpedia-api/controller"
	"debugpedia-api/db"
	"debugpedia-api/repository"
	"debugpedia-api/router"
	"debugpedia-api/usecase"
	"debugpedia-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	debugValidator := validator.NewDebugValidator()
	userRepository := repository.NewUserRepository(db)
	debugRepository := repository.NewDebugRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	debugUsecase := usecase.NewDebugUsecase(debugRepository, debugValidator)
	userController := controller.NewUserController(userUsecase)
	debugController := controller.NewDebugController(debugUsecase)
	e := router.NewRouter(userController, debugController)
	e.Logger.Fatal(e.Start(":8080"))
}

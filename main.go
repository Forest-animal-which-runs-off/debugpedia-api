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
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"fmt"
	"debugpedia-api/db"
	"debugpedia-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Debug{})
}
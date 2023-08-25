package main

import (
	"debugpedia-api/db"
	"debugpedia-api/model"
	"fmt"
)

func main(){
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	// dbに反映させたいモデル構造をアフィールドの値を0値でインスタンスし、アンバサンドで渡す。
	dbConn.AutoMigrate(&model.User{},&model.Debug{})
}
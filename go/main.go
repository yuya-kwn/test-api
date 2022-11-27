package main

import (
	"gin-test/go/database"
	"fmt"
	"gin-test/go/routes"
)

func main() {
	db := database.Connect()
	defer db.Close()

	err := db.Ping()
	if err != nil{
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功")
	}

	router := routes.NewRoutes()
	router.Run()
}

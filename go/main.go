package main

import (
	"gin-test/go/database"
	"gin-test/go/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.ConnectDatabase(1)

	router := routes.NewRoutes()
	router.Run()

}

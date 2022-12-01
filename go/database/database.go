package database

import (
	"database/sql"
	"fmt"
	"log"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase(tryCount int8) {
	dsn := "webuser:webpass@tcp(db:3306)/go_mysql8_development?charset=utf8mb4"

	db, err := sql.Open("mysql", dsn)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()

	if err != nil {
		if tryCount <= 5 {
			fmt.Println("データベース接続 - 再試行...")
			time.Sleep(time.Second * 3)
			ConnectDatabase(tryCount + 1)
		} else {
			log.Fatalln("データベース接続失敗")
		}
	} else {
		log.Println("データベース接続成功")
	}
}

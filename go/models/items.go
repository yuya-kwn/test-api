package models

import (

	_ "github.com/go-sql-driver/mysql"
)

type Items struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

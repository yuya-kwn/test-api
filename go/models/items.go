package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Items struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

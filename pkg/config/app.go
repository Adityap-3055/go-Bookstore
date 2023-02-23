package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// point of this file is to return a variable called db

var (
	db *gorm.DB
)

// connect function used to open a connection with our database
func Connect() {
	d, err := gorm.Open("mysql", "root:kinshu402@tcp/bookdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:root@/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
	log.Println("Database connected successfully")
}

func GetDB() *gorm.DB {

	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed")
	}
}

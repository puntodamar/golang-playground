package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "puntodamar:puntodamar@/golang_playground?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		log.Panicf( "Failed to open connection to database: %s", err.Error())
	}
}

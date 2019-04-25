package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	Db, e := gorm.Open("mysql", "puntodamar:puntodamar@/golang_playground?charset=utf8&parseTime=True&loc=Local")
	if e != nil{
		fmt.Println(e)
	}
}

package model

import (
	"github.com/jinzhu/gorm"
	"reflect"
	. "models/v1"
)

type ReturnValue struct {
	Model interface{}
}

func Exists(db *gorm.DB, model interface{}, column string, value string) bool {
	var data = ReturnValue{Model: model}
	var count int

	//db.Model(&model).Where(column + " = ?", value).Count(&count)

	switch (reflect.TypeOf(data.Model)){
	case reflect.TypeOf(User{}) :
		db.Model(User{}).Where(column + " = ?", value).Count(&count)
		break;
	}

	if count == 0{
		return false
	}
	return true
}


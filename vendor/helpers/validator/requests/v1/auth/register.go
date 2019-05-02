package auth

import (
	"fmt"
	"github.com/jinzhu/gorm"
	. "helpers/validator/error_formats"
	global_validator "helpers/validator/global"
	model_validator "helpers/validator/model"
	model "models/v1"
)

type RegisterForm struct {
	Name 		string `json:"name"`
	Username  	string `json:"username"`
	Email		string `json:"email"`
	Password 	string `json:"password"`
}


func (v *RegisterForm) Validate(db *gorm.DB) (bool, map[string][]ErrorFormat){

	qe := map[string][]ErrorFormat{}
	vu := validateUsername(db, v.Username)
	ve := validateEmail(db, v.Email)

	fmt.Println(len(vu))
	if len(vu) > 0 {
		qe["errors"] = append(qe["errors"], ErrorFormat{
			Field: "username", Messages: vu})
	}

	if len(ve) > 0 {
		qe["errors"] = append(qe["errors"], ErrorFormat{
			Field: "username", Messages: ve})
	}
	if(len(qe["errors"]) > 0){
		return false, qe
	} else {
		return true, qe
	}
}

func validateUsername(db *gorm.DB, username string) []interface{}{

	e := []interface{}{}

	// request exists
	if stat, msg := global_validator.Exists(username,"username"); stat == false{
		e = append(e, msg)
	}

	// username exists
	if stat, msg := model_validator.Exists(db, &model.User{}, "username", username); stat == true{
		e = append(e, msg)
	}

	// username length
	if stat,msg := global_validator.Between("username", len(username), 3,10); stat == false {
		e = append(e, msg)
	}

	return e
}

func validateEmail(db *gorm.DB, email string) []interface{}{

	e := []interface{}{}

	// request exists
	if stat, msg := global_validator.Exists(email,"email"); stat == false{
		e = append(e, msg)
	}

	// email valid
	if stat, msg := global_validator.ValidEmail(email); stat == false {
		e = append(e, msg)
	}

	// email exists
	if stat, msg := model_validator.Exists(db, &model.User{}, "email", email); stat == true{
		e = append(e, msg)
	}

	return e
}

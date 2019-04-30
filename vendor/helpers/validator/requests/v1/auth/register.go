package auth

import (
	"github.com/jinzhu/gorm"

	. "helpers/validator/error_formats"
	global_validator "helpers/validator/global"
	model_validator "helpers/validator/model"
	model "models/v1"
)

type RegisterForm struct {
	Name 		string `json:"name" validate:"required"`
	Username  	string `json:"username" validate:"required"`
	Email		string `json:"email" validate:"required,email"`
	Password 	string `json:"password" validate:"required"`
}


func (v *RegisterForm) Validate(db *gorm.DB, user *model.User) (bool, map[string][]ErrorFormat){

	qe := map[string][]ErrorFormat{
		"errors" : {
			ErrorFormat{
				Field	: "username",
				Messages: validateUsername(db, user.Username)},

			ErrorFormat{
				Field	: "email",
				Messages: 	validateEmail(db, user.Email)},
		},
	}

	if(len(qe) > 0){
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
	if stat, msg := model_validator.Exists(db, model.User{}, "username", username); stat == true{
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
	if stat, msg := global_validator.ValidEmail(email); stat == true{
		e = append(e, msg)
	}

	// email exists
	if stat, msg := model_validator.Exists(db, model.User{}, "email", email); stat == true{
		e = append(e, msg)
	}

	return e
}

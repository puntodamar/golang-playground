package auth

import (
	_ "helpers/validator"
)

type RegisterForm struct {
	Name 		string `json:"name" 	valid:"required"`
	Username  	string `json:"username" valid:"required"`
	Email		string `json:"email" 	valid:"email"`
	Password 	string `json:"password" valid:"required,length(6|20)"`
}


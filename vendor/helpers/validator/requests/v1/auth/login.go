package auth

//type LoginForm struct {
//	Username  	string `json:"username" binding:"ip"`
//	Password 	string `json:"password" validate:"required"`
//}

type LoginForm struct {
	Username  	string `json:"username"`
	Password 	string `json:"password"`
}

func (v *LoginForm) Validate() []string {

	errs := []string{}

	if len(v.Username) == 0 {
		errs = append(errs, "username is required")
	}

	if len(v.Password) == 0 {
		errs = append(errs, "password is required")
	}

	return errs
}
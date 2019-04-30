package auth

type LoginForm struct {
	Username  	string `json:"username"`
	Password 	string `json:"password"`
}

func (v *LoginForm) Validate(l string) []string {

	errs := []string{}

	if len(v.Username) < 1 {
		errs = append(errs, "username is required")
	}

	if len(v.Password) < 1 {
		errs = append(errs, "password is required")
	}

	return errs
}
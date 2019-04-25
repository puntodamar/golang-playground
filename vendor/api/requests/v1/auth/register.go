package auth

type RegisterForm struct {
	Name 		string `json:"name"`
	Username  	string `json:"username"`
	Email		string `json:"email"`
	Password 	string `json:"password"`
}

func (v *RegisterForm) Validate(language string) []string {

	errs := []string{}

	if len(v.Name) < 1 {
		errs = append(errs, "name is required")
	}

	if len(v.Username) < 1 {
		errs = append(errs, "username is required")
	}

	if len(v.Password) < 1 {
		errs = append(errs, "password is required")
	}

	if len(v.Email) < 1 {
		errs = append(errs, "email is required")
	}

	return errs
}
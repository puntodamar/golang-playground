package global

import (
	"regexp"
	"strconv"
)

func Between(col string, len int, min int, max int) (bool, string) {

	if len < min || len > max {
		return false, col + " should be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max)
	} else {
		return true, ""
	}
}

func ValidEmail(s string) (bool, string) {
	rgx := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if(rgx.MatchString(s)){
		return true, ""
	} else {
		return false, "invalid email"
	}
}

func Exists(s string, col string) (bool, string){
	if len(s) == 0 {
		return false, "field " + col + " is required"
	} else{
		return true,""
	}
}
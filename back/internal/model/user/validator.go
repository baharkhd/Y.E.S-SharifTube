package user

import "regexp"

func validateEmail(email string) bool {
	if len(email)==0 {
		return true
	}
	match, _ := regexp.MatchString("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)", email)
	return match
}

func validateBasicField(input string) bool {
	return len(input) < 255
}

func validate(username,name,email string)bool{
	return validateEmail(email) && validateBasicField(username) && validateBasicField(name)
}

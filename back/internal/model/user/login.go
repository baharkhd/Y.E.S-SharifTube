package user

import (
	"yes-sharifTube/graph/model"
	"yes-sharifTube/pkg/jwt"
)

func Login(username, password string) (string, error) {

	// retrieve user from data base
	// in case of error should return model.UserPassMissMatchException due to safety reasons
	target, err := Get(username)
	if err != nil {
		return "", model.UserPassMissMatchException{}
	}

	// check if the username and password matches
	if !target.Verify(password) {
		return "", model.UserPassMissMatchException{}
	}

	// generate new token
	token, err2 := jwt.GenerateToken(target.Name)
	if err2 != nil {
		return "", model.InternalServerException{}
	}
	return token, nil
}


package user

import (
"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) (string,error) {
	if hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost);
		err != nil {
		return "", err
	}else{
		return string(hash),nil
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *User) Verify(password string) bool {
	return checkPasswordHash(password, u.Password)
}

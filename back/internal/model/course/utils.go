package course

import "golang.org/x/crypto/bcrypt"

func hashToken(pwd []byte) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost);
		err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

func CheckTokenHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

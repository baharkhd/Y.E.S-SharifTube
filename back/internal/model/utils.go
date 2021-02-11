package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashToken(pwd []byte) (string, error) {
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

func ConvertStringsToObjectIDs(arr []string) ([]primitive.ObjectID, error) {
	var coIDs []primitive.ObjectID
	for _, cID := range arr {
		objID, err := primitive.ObjectIDFromHex(cID)
		if err != nil {
			return nil, err
		}
		coIDs = append(coIDs, objID)
	}
	return coIDs, nil
}

func ContainsInStringArray(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func RemoveFromStringArray(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func PtrTOStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

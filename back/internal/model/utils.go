package model

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
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

func IsSTREmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func WordCount(value string) int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return len(results)
}
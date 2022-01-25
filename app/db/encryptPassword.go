package db

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err

}

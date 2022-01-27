package db

import (
	"github.com/esmejia277/twittor/app/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckIfUserExists(email)

	if !found {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	error := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if error != nil {
		return user, false
	}
	return user, true

}

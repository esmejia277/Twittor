package router

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/esmejia277/twittor/app/db"
	"github.com/esmejia277/twittor/app/models"
	"github.com/pkg/errors"
)

var Email string
var IDUser string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	secretKey := []byte("topsecret")
	claims := &models.Claim{}
	splittedToken := strings.Split(token, "Bearer")
	if len(splittedToken) != 2 {
		return claims, false, "", errors.New("Invalid token")
	}
	token = strings.TrimSpace(splittedToken[1])
	validatedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return claims, false, "", errors.New("Invalid token")
	}
	_, found, _ := db.CheckIfUserExists(claims.Email)
	if found {
		Email = claims.Email
		IDUser = claims.ID.Hex()
		return claims, found, IDUser, nil
	}

	if !validatedToken.Valid {
		return claims, false, "", errors.New("Invalid token")
	}

	return claims, false, "", err
}

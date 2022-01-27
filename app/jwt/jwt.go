package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/esmejia277/twittor/app/models"
)

func GenerateJWT(model models.User) (string, error) {
	secretKey := []byte("topsecret")
	payload := jwt.MapClaims{
		"email":    model.Email,
		"name":     model.Name,
		"lastName": model.LastName,
		"birthday": model.Birthday,
		"_id":      model.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

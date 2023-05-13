package token

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "someSecret_addThisToConf"

func GenerateToken(username string, role string) (string, error) {
	tokenLifespan := 1 // in hrs

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	tokenG := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenG.SignedString([]byte(SECRET_KEY))
}

package token

import (
	"errors"
	"fmt"
	"strings"
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

/*func TokenValid(ctx *gin.Context, role string) error {
	tokenString := extractToken(ctx)
	token, err := parseToken(tokenString)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid && claims["role"] == role {
		return nil
	}

	return errors.New("token or role not valid")
}*/

func ExtractTokenUsername(tokenAll string) (string, error) {
	tokenString := extractToken(tokenAll)
	token, err := parseToken(tokenString)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprintf("%s", claims["username"]), nil
	}
	return "", nil
}

func extractToken(tokenString string) string {
	if len(strings.Split(tokenString, " ")) == 2 {
		return strings.Split(tokenString, " ")[1]
	}
	return ""
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token not signed properly")
		}
		return []byte(SECRET_KEY), nil
	})
	return token, err
}

/*func UserAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpGin := HTTP.Gin{Context: ctx}
		err := TokenValid(ctx, "USERROLE")
		if err != nil {
			ctx.Abort()
			httpGin.Unauthorized(nil)
			return
		}
		ctx.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpGin := HTTP.Gin{Context: ctx}
		err := TokenValid(ctx, "ADMINROLE")
		if err != nil {
			ctx.Abort()
			httpGin.Unauthorized(nil)
			return
		}
		ctx.Next()
	}
}*/

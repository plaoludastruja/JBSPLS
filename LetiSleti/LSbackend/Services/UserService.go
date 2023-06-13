package Services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
)

const SECRET_KEY = "someSecret_addThisToConf"

func GetAllUsers() []Models.User {
	return Repository.GetAllUsers()
}

func GetUserByEmail(email string) (Models.User, error) {
	return Repository.GetUserByEmail(email)
}

func DeleteUser(userId string) int64 {
	return Repository.DeleteUser(userId)
}

func GenerateApiKey(email string, isDurable string) {
	skitnicaUser, _ := Repository.GetSkitnicaUserByEmail(email)
	fmt.Println(skitnicaUser.Id)
	token, _ := GenerateToken(email, isDurable)
	_ = Repository.UpdateSkitnicaUserApiKey(token, skitnicaUser.Id.Hex())
}

func GenerateToken(email string, isDurable string) (string, error) {
	var tokenLifespan int // in hrs
	if isDurable == "true" {
		tokenLifespan = 240 // 10 days
	} else {
		tokenLifespan = 876000 //100 years
	}
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token.SignedString([]byte(SECRET_KEY)))
	return token.SignedString([]byte(SECRET_KEY))
}

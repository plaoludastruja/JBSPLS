package Services

import (
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/Token"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user Models.User) bool {
	HashPassword(&user)
	return Repository.CreateUser(user)
}

func LoginCheck(email string, password string) (string, error) {
	u, errr := Repository.GetUserByEmail(email)
	if errr != nil {
		return "", errr
	}

	err := CheckPasswordHash(password, u.Password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := Token.GenerateToken(u.Email, u.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func HashPassword(u *Models.User) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

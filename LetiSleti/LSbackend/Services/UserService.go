package Services

import (
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user Models.User) bool {
	HashPassword(&user)
	return Repository.CreateUser(user)
}

func GetAllUsers() []Models.User {
	return Repository.GetAllUsers()
}

func HashPassword(u *Models.User) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
	return nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

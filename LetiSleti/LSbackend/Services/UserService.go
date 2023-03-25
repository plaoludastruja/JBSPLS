package Services

import (
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
)

func GetAllUsers() []Models.User {
	return Repository.GetAllUsers()
}

func DeleteUser(userId string) int64 {
	return Repository.DeleteUser(userId)
}

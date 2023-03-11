package Services

import (
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
)

func RegisterUser(user Models.User) bool {
	if Repository.CreateUser(user) {
		return true
	}
	return false
}

package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/HTTP"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Services"
)

func GetAllUsers(ctx *gin.Context) {
	fmt.Println("GetAllUsers")
	httpGin := HTTP.Gin{Context: ctx}
	users := Services.GetAllUsers()
	httpGin.OK(users)
}

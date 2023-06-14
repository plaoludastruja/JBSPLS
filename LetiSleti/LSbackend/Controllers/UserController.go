package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/HTTP"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Services"
)

func GetAllUsers(ctx *gin.Context) {
	fmt.Println("GetAllUsers")
	httpGin := HTTP.Gin{Context: ctx}
	users := Services.GetAllUsers()
	httpGin.OK(users)
}

func GetUserByEmail(ctx *gin.Context) {
	fmt.Println("GetUserByEmail")
	httpGin := HTTP.Gin{Context: ctx}
	user := Models.User{}
	email := ctx.Param("email")
	fmt.Println(email)
	user, _ = Services.GetUserByEmail(email)
	httpGin.OK(user)
}

func DeleteUser(ctx *gin.Context) {
	fmt.Println("DeleteUser")
	userId := ctx.Param("userId")
	httpGin := HTTP.Gin{Context: ctx}
	users := Services.DeleteUser(userId)
	httpGin.OK(users)
}

func GenerateApiKey(ctx *gin.Context) {
	fmt.Println("GenerateApiKey")
	email := ctx.Param("email")
	isDurable := ctx.Param("isDurable")
	fmt.Println(email)
	httpGin := HTTP.Gin{Context: ctx}
	Services.GenerateApiKey(email, isDurable)
	httpGin.OK("Ok")
}

package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/HTTP"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Services"
)

func RegisterUser(ctx *gin.Context) {
	fmt.Println("RegisterUser")
	httpGin := HTTP.Gin{Context: ctx}
	user := Models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		httpGin.BadRequest(err.Error())
		return
	}

	if !Services.RegisterUser(user) {
		httpGin.BadRequest(user)
		return
	}

	httpGin.Created(user)

}

func GetAllUsers(ctx *gin.Context) {
	fmt.Println("GetAllUsers")
	httpGin := HTTP.Gin{Context: ctx}
	users := []Models.User{}

	users = Services.GetAllUsers()

	httpGin.OK(users)

}

func Hello(ctx *gin.Context) {
	httpGin := HTTP.Gin{Context: ctx}
	httpGin.OK("Hello")
}

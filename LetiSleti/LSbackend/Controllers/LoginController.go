package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/HTTP"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models/DTO"
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

func LoginUser(ctx *gin.Context) {
	fmt.Println("LoginUser")
	httpGin := HTTP.Gin{Context: ctx}
	user := DTO.UserDTO{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		httpGin.BadRequest(err.Error())
		return
	}
	token, err := Services.LoginCheck(user.Email, user.Password)
	if err != nil {
		httpGin.BadRequest(err.Error())
		return
	}
	httpGin.Created(token)
}

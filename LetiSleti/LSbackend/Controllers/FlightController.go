package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/HTTP"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Services"
)

func RegisterFlight(ctx *gin.Context) {
	fmt.Println("RegisterFlight")
	httpGin := HTTP.Gin{Context: ctx}
	flight := Models.Flight{}
	if err := ctx.ShouldBindJSON(&flight); err != nil {
		httpGin.BadRequest(err.Error())
		return
	}
	if !Services.AddFlight(flight) {
		httpGin.BadRequest(flight)
		return
	}
	httpGin.Created(flight)
}

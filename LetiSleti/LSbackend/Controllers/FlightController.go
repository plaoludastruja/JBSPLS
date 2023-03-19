package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/HTTP"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models/DTO"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Services"
)

func RegisterFlight(ctx *gin.Context) {
	fmt.Println("RegisterFlight")
	httpGin := HTTP.Gin{Context: ctx}
	flight := Models.Flight{}
	if err := ctx.ShouldBindJSON(&flight); err != nil {
		fmt.Println("Tu je")
		fmt.Println(&flight)
		httpGin.BadRequest(err.Error())
		return
	}
	if !Services.AddFlight(flight) {
		httpGin.BadRequest(flight)
		return
	}
	httpGin.Created(flight)
}

func GetAllFlights(ctx *gin.Context) {
	fmt.Println("GetAllFlights")
	httpGin := HTTP.Gin{Context: ctx}
	flights := Services.GetAllFlights()
	httpGin.OK(flights)
}

func SearchFlights(ctx *gin.Context) {
	fmt.Println("SearchFlights")
	httpGin := HTTP.Gin{Context: ctx}
	searchCriteria := DTO.SearchDTO{}

	if err := ctx.ShouldBindJSON(&searchCriteria); err != nil {
		fmt.Println("Cannot bind json")
		fmt.Println(&searchCriteria)
		httpGin.BadRequest(err.Error())
		return
	}

	flights := Services.SearchFlights(searchCriteria)
	httpGin.OK(flights)
}

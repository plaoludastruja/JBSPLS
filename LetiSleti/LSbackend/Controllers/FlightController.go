package Controllers

import (
	"fmt"
	"strconv"

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
	fmt.Println("search criteria controller:")
	fmt.Println(searchCriteria.StartPlace)
	fmt.Println(searchCriteria.EndPlace)
	fmt.Println(searchCriteria.Date)
	flights := Services.SearchFlights(searchCriteria)
	httpGin.OK(flights)
}

func DeleteFlight(ctx *gin.Context) {
	fmt.Println("Delete flight")
	httpGin := HTTP.Gin{Context: ctx}
	flightId := ctx.Param("flightId")
	fmt.Println(flightId)
	ret := Services.DeleteFlight(flightId)

	httpGin.OK(ret)
}

func ChangePlacesLeft(ctx *gin.Context) {
	fmt.Println("ChangePlacesLeft")
	httpGin := HTTP.Gin{Context: ctx}

	flightId := ctx.Param("flightId")
	count := ctx.Param("count")
	fmt.Println(count)
	fmt.Println(flightId)
	countConverted, _ := strconv.Atoi(count)
	Services.ChangePlacesLeft(flightId, countConverted)

	httpGin.OK(flightId)
}

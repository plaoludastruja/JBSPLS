package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/HTTP"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/Token"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models/DTO"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Services"
)

func CreateTicket(ctx *gin.Context) {
	fmt.Println("Creating ticket")
	httpGin := HTTP.Gin{Context: ctx}
	ticket := Models.Ticket{}
	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		fmt.Println(&ticket)
		httpGin.BadRequest(err.Error())
		return
	}
	if !Services.CreateTicket(ticket) {
		httpGin.BadRequest(ticket)
		return
	}
	httpGin.Created(ticket)
}

func GetAllTickets(ctx *gin.Context) {
	fmt.Println("GetAllTickets")
	httpGin := HTTP.Gin{Context: ctx}
	email, _ := Token.ExtractTokenEmail(ctx)
	fmt.Println(email)
	tickets := Services.GetAllTickets(email)
	httpGin.OK(tickets)
}

func CreateTicketFromSkitnica(ctx *gin.Context) {
	fmt.Println("Creating ticket from skitnica")
	httpGin := HTTP.Gin{Context: ctx}
	ticketSkitnicaDto := DTO.TicketSkitnicaDTO{}
	if err := ctx.ShouldBindJSON(&ticketSkitnicaDto); err != nil {
		fmt.Println(&ticketSkitnicaDto)
		httpGin.BadRequest(err.Error())
		return
	}
	if !Services.CreateTicketFromSkitnica(ticketSkitnicaDto) {
		httpGin.BadRequest(ticketSkitnicaDto)
		return
	}
	httpGin.Created(ticketSkitnicaDto)
}

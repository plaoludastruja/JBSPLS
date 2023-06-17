package Services

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/Token"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models/DTO"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
)

func CreateTicket(ticket Models.Ticket) bool {
	return Repository.CreateTicket(ticket)
}

func GetAllTickets(email string) []Models.Ticket {
	return Repository.GetAllTickets(email)
}

func CreateTicketFromSkitnica(ticketSkitnicaDto DTO.TicketSkitnicaDTO) bool {
	exp, _ := Token.ExtractApiKeyExp(ticketSkitnicaDto.ApiKey)
	split := strings.Split(exp, "-")
	year, _ := strconv.Atoi(split[0])
	month, _ := strconv.Atoi(split[1])
	day, _ := strconv.Atoi(split[2])
	end := time.Date(year, time.Month(month), day, 0, 0, 0, 999999999, time.UTC)
	fmt.Println("exp")
	fmt.Println(exp)
	if end.Before(time.Now()) {
		fmt.Println("istekao")
		return false
	}
	email, _ := Token.ExtractApiKeyEmail(ticketSkitnicaDto.ApiKey)

	user, _ := Repository.GetUserByEmail(email)
	ticket := Models.Ticket{Email: user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Start:      ticketSkitnicaDto.Start,
		StartPlace: ticketSkitnicaDto.StartPlace,
		End:        ticketSkitnicaDto.End,
		EndPlace:   ticketSkitnicaDto.EndPlace,
		Price:      ticketSkitnicaDto.Price,
		Count:      ticketSkitnicaDto.Count,
		FlightId:   ticketSkitnicaDto.FlightId}
	return Repository.CreateTicket(ticket)
}

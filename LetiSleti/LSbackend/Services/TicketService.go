package Services

import (
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
)

func CreateTicket(ticket Models.Ticket) bool {
	return Repository.CreateTicket(ticket)
}

func GetAllTickets(email string) []Models.Ticket {
	return Repository.GetAllTickets(email)
}

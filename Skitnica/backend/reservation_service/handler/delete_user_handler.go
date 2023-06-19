package handler

import (
	events "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/create_order"
	saga "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/messaging"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/service"
)

type DeleteUserCommandHandler struct {
	reservationService *service.ReservationService
	replyPublisher     saga.Publisher
	commandSubscriber  saga.Subscriber
}

func NewDeleteUserCommandHandler(reservationService *service.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		reservationService: reservationService,
		replyPublisher:     publisher,
		commandSubscriber:  subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteUserCommandHandler) handle(command *events.DeleteUserCommand) {

	reply := events.DeleteUserReply{User: command.User}

	switch command.Type {
	case events.CheckReservations:
		res, err := handler.reservationService.CheckReservations(command.User)
		if err != nil {
			return
		}
		if res {
			reply.Type = events.CanDelete
		} else {
			reply.Type = events.CanNotDelete
		}

	default:
		reply.Type = events.UnknownRep
	}

	if reply.Type != events.UnknownRep {
		_ = handler.replyPublisher.Publish(reply)
	}
}

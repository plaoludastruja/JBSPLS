package handler

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/service"
	events "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/create_order"
	saga "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/messaging"
)

type DeleteUserCommandHandler struct {
	accomodationService *service.AccomodationService
	replyPublisher      saga.Publisher
	commandSubscriber   saga.Subscriber
}

func NewDeleteUserCommandHandler(accomodationService *service.AccomodationService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		accomodationService: accomodationService,
		replyPublisher:      publisher,
		commandSubscriber:   subscriber,
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
	case events.DeleteAccomodations:
		res, err := handler.accomodationService.GetAll()

		if err != nil {
			reply.Type = events.UnknownRep
			return
		}
		for _, accomodation := range res {
			if accomodation.HostUsername == command.User.Username {
				er := handler.accomodationService.Delete(accomodation.Id)
				if er != nil {
					reply.Type = events.UnknownRep
					break
				}
			}
		}

		reply.Type = events.AccomodationsDeleted

	default:
		reply.Type = events.UnknownRep
	}

	if reply.Type != events.UnknownRep {
		_ = handler.replyPublisher.Publish(reply)
	}
}

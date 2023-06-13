package handler

import (
	events "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/create_order"
	saga "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/messaging"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/service"
)

type DeleteUserCommandHandler struct {
	userService       *service.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeleteUserCommandHandler(userService *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		userService:       userService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
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
	case events.DeleteUser:
		err := handler.userService.DelUser(command.User)
		if err != nil {
			reply.Type = events.UnknownRep
		}
		if command.User.Role == "USER" {
			reply.Type = events.DeletedGuest
		} else {
			reply.Type = events.DeletedHost
		}

	default:
		reply.Type = events.UnknownRep
	}

	if reply.Type != events.UnknownRep {
		_ = handler.replyPublisher.Publish(reply)
	}
}

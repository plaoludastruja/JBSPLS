package service

import (
	events "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/create_order"
	saga "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/messaging"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/domain"
)

type DeleteUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewDeleteUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserOrchestrator, error) {
	o := &DeleteUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *DeleteUserOrchestrator) Start(user *domain.User) error {
	event := &events.DeleteUserCommand{
		Type: events.CheckReservations,
		User: events.User{
			Id:        user.Id,
			Username:  user.Username,
			Password:  user.Password,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			Address:   user.Address,
		},
	}

	return o.commandPublisher.Publish(event)
}

func (o *DeleteUserOrchestrator) handle(reply *events.DeleteUserReply) {
	command := events.DeleteUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.Unknown {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *DeleteUserOrchestrator) nextCommandType(reply events.DeleteUserReplyType) events.DeleteUserCommandType {
	switch reply {
	case events.CanDelete:
		return events.DeleteUser
	case events.DeletedHost:
		return events.DeleteAccomodations
	default:
		return events.Unknown
	}
}

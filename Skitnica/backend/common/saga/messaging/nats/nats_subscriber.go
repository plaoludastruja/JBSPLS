package nats

import (
	saga "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/messaging"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	conn       *nats.EncodedConn
	subject    string
	queueGroup string
}

func NewNATSSubscriber(host, port, user, password, subject, queueGroup string) (saga.Subscriber, error) {
	conn, err := getConnection(host, port, user, password)
	if err != nil {
		return nil, err
	}
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		conn:       encConn,
		subject:    subject,
		queueGroup: queueGroup,
	}, nil
}

func (s *Subscriber) Subscribe(handler interface{}) error {
	_, err := s.conn.QueueSubscribe(s.subject, s.queueGroup, handler)
	if err != nil {
		return err
	}
	return nil
}

/*func (s *Subscriber) SubscribeNotification() (string, error) {
	_, err := s.conn.QueueSubscribe(s.subject, s.queueGroup, func(message *nats.Msg) (string, error) {
		fmt.Printf("RECEIVED MESSAGE: %s\n", string(message.Data))
		return string(message.Data), nil
	})
	if err != nil {
		return "", err
	}
	return "", nil
}*/

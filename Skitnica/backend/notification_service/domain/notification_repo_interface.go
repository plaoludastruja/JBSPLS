package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type INotificationRepo interface {
	Get(id primitive.ObjectID) (*Notification, error)
	GetBySender(sender string) ([]*Notification, error)
	GetByReceiver(receiver string) ([]*Notification, error)
	GetAll() ([]*Notification, error)
	Insert(notification *Notification) error
	Edit(usenotificationr *Notification) error
	DeleteAll()
	Delete(id primitive.ObjectID) error
	InsertNotificationFilters(notificationFilter *NotificationFilter) error
	EditNotificationFilter(notificationFilter *NotificationFilter) error
	GetNotificationFilterByUsername(username string) (*NotificationFilter, error)
}

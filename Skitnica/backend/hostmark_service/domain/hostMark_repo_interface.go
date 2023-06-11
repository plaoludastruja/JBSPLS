package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IHostMarkRepo interface {
	Get(id primitive.ObjectID) (*HostMark, error)
	GetAll() ([]*HostMark, error)
	Insert(user *HostMark) error
	Edit(user *HostMark) error
	DeleteAll()
	Delete(id primitive.ObjectID) error
	GetByUsername(username string, hostUsername string) ([]*HostMark, error)
	GetAllByHostUsername(hostUsername string) ([]*HostMark, error)
}

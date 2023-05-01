package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepo interface {
	Get(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) error
	DeleteAll()
}

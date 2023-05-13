package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepo interface {
	Get(id primitive.ObjectID) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) error
	Edit(user *User) error
	DeleteAll()
	Delete(id primitive.ObjectID) error
}

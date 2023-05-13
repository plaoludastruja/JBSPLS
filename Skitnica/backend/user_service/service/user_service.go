package service

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.IUserRepo
}

func NewUserService(store domain.IUserRepo) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetByUsername(username string) (*domain.User, error) {
	return service.store.GetByUsername(username)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Insert(user domain.User) error {
	return service.store.Insert(&user)
}

func (service *UserService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *UserService) Edit(user domain.User) error {
	return service.store.Edit(&user)
}

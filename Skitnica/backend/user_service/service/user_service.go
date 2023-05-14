package service

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/domain"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
	//hashPassword(&user)
	_, err := service.GetByUsername(user.Username)
	if err == nil {
		return err
	}
	return service.store.Insert(&user)
}

func (service *UserService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *UserService) Edit(user domain.User) error {
	return service.store.Edit(&user)
}

func (service *UserService) Login(username string, password string) (string, error) {
	user, err := service.GetByUsername(username)
	if err != nil {
		return "", err
	}

	if password != user.Password {
		return "", bcrypt.ErrMismatchedHashAndPassword
	}

	/*errr := checkPasswordHash(password, user.Password)
	if errr == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}*/

	token, err := token.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func hashPassword(u *domain.User) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
}

func checkPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

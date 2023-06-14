package service

import (
	"context"
	"fmt"
	"log"

	reservationProto "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
	events "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/create_order"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/domain"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	store        domain.IUserRepo
	orchestrator *DeleteUserOrchestrator
}

func NewUserService(store domain.IUserRepo, orchestrator *DeleteUserOrchestrator) *UserService {
	return &UserService{
		store:        store,
		orchestrator: orchestrator,
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

	user, err := service.Get(id)

	if err != nil {
		return err
	}
	err = service.orchestrator.Start(user)

	if err != nil {
		return err
	}
	return nil

	/*user, err := service.Get(id)
	fmt.Println("13213")
	if err != nil {
		return err
	}
	if user.Role == "USER" {
		fmt.Println("1")
		return service.deleteUser(user)
	} else {
		fmt.Println("2")
		return service.deleteHost(user)
	}
	//return service.store.Delete(id)*/
}

func (service *UserService) deleteUser(user *domain.User) error {
	reservationEndpoint := fmt.Sprintf("%s:%s", "reservation_service", "8000")
	fmt.Println("3")
	conn, err := grpc.Dial(reservationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("4")
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	fmt.Println("5")
	reservationClient := reservationProto.NewReservationServiceClient(conn)
	fmt.Println("6")
	reservations, err := reservationClient.GetAllRes(context.TODO(), &reservationProto.GetAllReq{})
	fmt.Println("7")
	fmt.Println(len(reservations.Reservations))
	for _, reservation := range reservations.Reservations {
		fmt.Println("8")
		if reservation.Username == user.Username {
			fmt.Println("9")
			return fmt.Errorf("guest has a reservation in a future")
		}
	}
	fmt.Println("10")
	return service.store.Delete(user.Id)
}

func (service *UserService) deleteHost(user *domain.User) error {
	reservationEndpoint := fmt.Sprintf("%s:%s", "reservation_service", "8000")
	conn, err := grpc.Dial(reservationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	reservationClient := reservationProto.NewReservationServiceClient(conn)
	reservations, err := reservationClient.GetAllRes(context.TODO(), &reservationProto.GetAllReq{})

	for _, reservation := range reservations.Reservations {
		if reservation.HostUsername == user.Username {
			return fmt.Errorf("guest has a reservation in a future")
		}
	}

	return service.store.Delete(user.Id)
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

func (service *UserService) GetHosts() ([]*domain.User, error) {
	return service.store.GetHosts()
}

func (service *UserService) DelUser(user events.User) error {
	return service.store.Delete(user.Id)
}

package service

import (
	"context"
	"fmt"
	"log"
	"time"

	notificationProto "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	events "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/create_order"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ReservationService struct {
	store domain.IReservationRepo
}

func NewReservationService(store domain.IReservationRepo) *ReservationService {
	return &ReservationService{
		store: store,
	}
}

func (service *ReservationService) Get(id primitive.ObjectID) (*domain.Reservation, error) {
	return service.store.Get(id)
}

func (service *ReservationService) GetAll() ([]*domain.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) Insert(reservation domain.Reservation) error {
	return service.store.Insert(&reservation)
}

func (service *ReservationService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *ReservationService) Edit(reservation domain.Reservation) error {
	return service.store.Edit(&reservation)
}

func (service *ReservationService) Search(startDay string, startMonth string, startYear string, endDay string, endMonth string, endYear string) ([]*domain.Reservation, error) {
	return service.store.Search(startDay, startMonth, startYear, endDay, endMonth, endYear)
}
func (service *ReservationService) Check(dateRange domain.DateRange) ([]*domain.Reservation, error) {
	return service.store.Check(&dateRange)
}

func (service *ReservationService) GetAllPending(hostUsername string) ([]*domain.Reservation, error) {
	return service.store.GetAllPending(hostUsername)
}

func (service *ReservationService) GetCanceledForUser(username string) (int32, error) {
	list, _ := service.store.GetCanceledForUser(username)
	num := 0
	for _, reservation := range list {
		fmt.Println(reservation.Username)
		num = num + 1
	}
	return int32(num), nil
}

func (service *ReservationService) ApproveReservation(reservationDto domain.ReservationDto) {
	service.store.ApproveReservation(&reservationDto)
	service.store.RejectOverlapsed(&reservationDto)
	service.sendNotification(reservationDto, "prihvaćena. ")
}

func (service *ReservationService) RejectReservation(reservationDto domain.ReservationDto) {
	service.store.RejectReservation(&reservationDto)

	service.sendNotification(reservationDto, "odbijena. ")
}

func (service *ReservationService) GetForGuest(username string) []string {
	return service.store.GetForGuest(username)
}

func (service *ReservationService) IsBestHostCheck(hostUsername string) bool {
	reservations, _ := service.GetAllByHostUsername(hostUsername)
	numOfReservations := len(reservations)
	fmt.Println("num", numOfReservations)
	if numOfReservations < 5 {
		return false
	}

	canceledReservations, _ := service.GetAllCanceledByHostUsername(hostUsername)
	numOfCanceledReservations := len(canceledReservations)
	fmt.Println("canceled", numOfCanceledReservations)
	if float64(numOfCanceledReservations)/float64(numOfReservations) > 0.05 {
		return false
	}

	daysSum := 0
	for _, reservation := range reservations {
		duration := reservation.EndDate.Sub(reservation.StartDate)
		days := duration / (24 * time.Hour)
		daysSum = daysSum + int(days)
		fmt.Println("", days)
	}
	fmt.Println("sum", daysSum)
	if daysSum < 50 {
		return false
	}

	return true
}

func (service *ReservationService) GetAllByHostUsername(hostUsername string) ([]*domain.Reservation, error) {
	return service.store.GetAllByHostUsername(hostUsername)
}

func (service *ReservationService) GetAllCanceledByHostUsername(hostUsername string) ([]*domain.Reservation, error) {
	return service.store.GetAllCanceledByHostUsername(hostUsername)
}

func (service *ReservationService) CheckReservations(user events.User) (bool, error) {
	if user.Role == "USER" {
		return service.CheckReservationsForGuest(user)
	} else {
		return service.CheckReservationsForHost(user)
	}
}

func (service *ReservationService) CheckReservationsForGuest(user events.User) (bool, error) {

	reservations, _ := service.GetAll()

	for _, reservation := range reservations {
		if reservation.Username == user.Username {
			return false, nil
		}
	}

	return true, nil
}

func (service *ReservationService) CheckReservationsForHost(user events.User) (bool, error) {
	reservations, _ := service.GetAll()

	for _, reservation := range reservations {
		if reservation.HostUsername == user.Username {
			return false, nil
		}
	}

	return true, nil
}

func (service *ReservationService) sendNotification(reservationDto domain.ReservationDto, reservationStatus string) {
	// ako se ocjeni host, salje se notifikacija na notifiaction_service
	notificationEndpoint := fmt.Sprintf("%s:%s", "notification_service", "8000")
	conn, err := grpc.Dial(notificationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	notificationClient := notificationProto.NewNotificationServiceClient(conn)

	notificationPb := notificationProto.Notification{
		Id:       primitive.NewObjectID().Hex(),
		Receiver: reservationDto.Username,
		Sender:   reservationDto.AccomodationId,
		Subject:  "reservation",
		Message:  " Vaša rezervacija za datum " + reservationDto.StartDate.Format("01.02.2006.") + " - " + reservationDto.EndDate.Format("01.02.2006.") + " je " + reservationStatus,
		IsRead:   "false",
	}

	notification, err := notificationClient.CreateNotification(context.TODO(), &notificationProto.CreateNotificationRequest{Notification: &notificationPb})
	fmt.Println(notification.Notification)
}

package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	notificationProto "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DATABASE   = "reservationsDB"
	COLLECTION = "reservations"
)

type ReservationRepo struct {
	reservations *mongo.Collection
}

func NewReservationRepo(client *mongo.Client) domain.IReservationRepo {
	reservationsCollection := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationRepo{
		reservations: reservationsCollection,
	}
}

func (store *ReservationRepo) Get(id primitive.ObjectID) (*domain.Reservation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *ReservationRepo) GetAll() ([]*domain.Reservation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *ReservationRepo) Insert(reservation *domain.Reservation) error {
	result, err := store.reservations.InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReservationRepo) Edit(reservation *domain.Reservation) error {
	filter := bson.M{"_id": reservation.Id}
	fmt.Println(reservation.StartDate.Compare(time.Now().Add(time.Hour * 24)))
	fmt.Println(reservation.StartDate)
	fmt.Println(time.Now().Add(time.Hour * 24))
	if reservation.StartDate.Compare(time.Now().Add(time.Hour*24)) == 1 && reservation.Status == "APPROVED" {
		update := bson.M{"$set": bson.M{
			"accomodationId": reservation.AccomodationId,
			"username":       reservation.Username,
			"startDate":      reservation.StartDate,
			"endDate":        reservation.EndDate,
			"guestNumber":    reservation.GuestNumber,
			"status":         "CANCELED",
		}}
		_, err := store.reservations.UpdateOne(context.TODO(), filter, update)

		// ako se otkaze rezevacija, salje se notifikacija na notifiaction_service
		if err == nil {
			notificationEndpoint := fmt.Sprintf("%s:%s", "notification_service", "8000")
			conn, err := grpc.Dial(notificationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
			}
			notificationClient := notificationProto.NewNotificationServiceClient(conn)

			notificationPb := notificationProto.Notification{
				Id:       primitive.NewObjectID().Hex(),
				Receiver: reservation.HostUsername,
				Sender:   reservation.Username,
				Subject:  "CANCEL_RESERVATION",
				Message:  reservation.Username + " je otkazao rezevaciju kod " + reservation.HostUsername,
				IsRead:   "false",
			}

			notification, err := notificationClient.CreateNotification(context.TODO(), &notificationProto.CreateNotificationRequest{Notification: &notificationPb})
			fmt.Println(notification.Notification)
		}
		return err
	}
	return nil
}

func (store *ReservationRepo) DeleteAll() {
	store.reservations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *ReservationRepo) Delete(id primitive.ObjectID) error {
	reservation, _ := store.Get(id)
	if reservation.Status == "PENDING" {
		res, err := store.reservations.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
		fmt.Printf("deleted %v documents\n", res.DeletedCount)
		return err
	}
	return nil
}

func (store *ReservationRepo) filter(filter interface{}) ([]*domain.Reservation, error) {
	cursor, err := store.reservations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *ReservationRepo) filterOne(filter interface{}) (reservation *domain.Reservation, err error) {
	result := store.reservations.FindOne(context.TODO(), filter)
	err = result.Decode(&reservation)
	return
}

func decode(cursor *mongo.Cursor) (reservations []*domain.Reservation, err error) {
	for cursor.Next(context.TODO()) {
		var reservation domain.Reservation
		err = cursor.Decode(&reservation)
		if err != nil {
			return
		}
		reservations = append(reservations, &reservation)
	}
	err = cursor.Err()
	return
}

func (store *ReservationRepo) Search(startDay string, startMonth string, startYear string, endDay string, endMonth string, endYear string) ([]*domain.Reservation, error) {

	startYearInt, _ := strconv.Atoi(startYear)
	startMonthInt, _ := strconv.Atoi(startMonth)
	startDayInt, _ := strconv.Atoi(startDay)

	endYearInt, _ := strconv.Atoi(endYear)
	endMonthInt, _ := strconv.Atoi(endMonth)
	endDayInt, _ := strconv.Atoi(endDay)

	start := time.Date(startYearInt, time.Month(startMonthInt), startDayInt, 00, 00, 00, 999999999, time.UTC)
	end := time.Date(endYearInt, time.Month(endMonthInt), endDayInt, 00, 00, 00, 999999999, time.UTC)
	filter1 := bson.M{"startDate": bson.M{"$gte": start, "$lt": end}, "status": "APPROVED"}
	res1, _ := store.filter(filter1)

	filter2 := bson.M{"startDate": bson.M{"$lt": start}, "endDate": bson.M{"$gte": start}, "status": "APPROVED"}
	res2, _ := store.filter(filter2)

	res1 = append(res1, res2...)
	return res1, nil
}

func (store *ReservationRepo) Check(dateRange *domain.DateRange) ([]*domain.Reservation, error) {
	filter1 := bson.M{"startDate": bson.M{"$gte": dateRange.StartDate, "$lte": dateRange.EndDate}, "accomodationId": dateRange.AccomodationId, "status": "APPROVED"}
	res1, _ := store.filter(filter1)

	filter2 := bson.M{"endDate": bson.M{"$gte": dateRange.StartDate, "$lte": dateRange.EndDate}, "accomodationId": dateRange.AccomodationId, "status": "APPROVED"}
	res2, _ := store.filter(filter2)

	res1 = append(res1, res2...)
	return res1, nil
}

func (store *ReservationRepo) GetAllPending(hostUsername string) ([]*domain.Reservation, error) {
	filter := bson.M{"hostUsername": hostUsername, "status": "PENDING"}
	return store.filter(filter)
}

func (store *ReservationRepo) ApproveReservation(reservationDto *domain.ReservationDto) error {
	filter := bson.M{"_id": reservationDto.Id}
	fmt.Println(reservationDto.StartDate.Compare(time.Now().Add(time.Hour * 24)))
	fmt.Println(reservationDto.StartDate)
	fmt.Println(time.Now().Add(time.Hour * 24))

	update := bson.M{"$set": bson.M{
		"status": "APPROVED",
	}}
	_, err := store.reservations.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *ReservationRepo) RejectReservation(reservationDto *domain.ReservationDto) error {
	filter := bson.M{"_id": reservationDto.Id}
	fmt.Println(reservationDto.StartDate.Compare(time.Now().Add(time.Hour * 24)))
	fmt.Println(reservationDto.StartDate)
	fmt.Println(time.Now().Add(time.Hour * 24))
	if reservationDto.StartDate.Compare(time.Now().Add(time.Hour*24)) == 1 {
		update := bson.M{"$set": bson.M{
			"status": "REJECTED",
		}}
		_, err := store.reservations.UpdateOne(context.TODO(), filter, update)

		return err
	}
	return nil
}

func (store *ReservationRepo) GetCanceledForUser(username string) ([]*domain.Reservation, error) {
	filter := bson.M{"username": username, "status": "CANCELED"}
	return store.filter(filter)
}

func (store *ReservationRepo) RejectOverlapsed(reservationDto *domain.ReservationDto) error {

	update := bson.M{"$set": bson.M{
		"status": "REJECTED",
	}}

	filter1 := bson.M{"accomodationId": reservationDto.AccomodationId, "startDate": bson.M{"$lte": reservationDto.StartDate}, "endDate": bson.M{"$gte": reservationDto.StartDate}, "username": bson.M{"$ne": reservationDto.Username}}
	_, err := store.reservations.UpdateMany(context.TODO(), filter1, update)

	if err != nil {
		return err
	}

	filter2 := bson.M{"accomodationId": reservationDto.AccomodationId, "startDate": bson.M{"$lte": reservationDto.EndDate}, "endDate": bson.M{"$gte": reservationDto.EndDate}, "username": bson.M{"$ne": reservationDto.Username}}
	_, err2 := store.reservations.UpdateMany(context.TODO(), filter2, update)

	if err2 != nil {
		return err
	}

	filter3 := bson.M{"accomodationId": reservationDto.AccomodationId, "startDate": bson.M{"$gte": reservationDto.StartDate}, "endDate": bson.M{"$lte": reservationDto.EndDate}, "username": bson.M{"$ne": reservationDto.Username}}
	_, err3 := store.reservations.UpdateMany(context.TODO(), filter3, update)

	if err3 != nil {
		return err
	}

	return nil
}

func (store *ReservationRepo) GetForGuest(username string) []string {
	filter := bson.M{"username": username, "status": "APPROVED", "endDate": bson.M{"$lte": time.Now()}}
	reservations, _ := store.filter(filter)
	usernames := []string{}
	for i := 0; i < len(reservations); i++ {
		if !contains(usernames, reservations[i].HostUsername) {
			usernames = append(usernames, reservations[i].HostUsername)
		}
	}
	return usernames
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (store *ReservationRepo) GetAllByHostUsername(hostUsername string) ([]*domain.Reservation, error) {
	filter := bson.M{"hostUsername": hostUsername}
	return store.filter(filter)
}

func (store *ReservationRepo) GetAllCanceledByHostUsername(hostUsername string) ([]*domain.Reservation, error) {
	filter := bson.M{"hostUsername": hostUsername, "status": "CANCELED"}
	return store.filter(filter)
}

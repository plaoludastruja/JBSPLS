package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

		return err
	}
	return nil
}

func (store *ReservationRepo) DeleteAll() {
	store.reservations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *ReservationRepo) Delete(id primitive.ObjectID) error {
	reservation, _ := store.Get(id)
	if reservation.Status == "PENDING" || (reservation.StartDate.Compare(time.Now().Add(time.Hour*48)) == -1 && reservation.Status == "APPROVED") {
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
	filter1 := bson.M{"startDate": bson.M{"$gte": dateRange.StartDate, "$lte": dateRange.EndDate}}
	res1, _ := store.filter(filter1)

	filter2 := bson.M{"endDate": bson.M{"$gte": dateRange.StartDate, "$lte": dateRange.EndDate}}
	res2, _ := store.filter(filter2)

	res1 = append(res1, res2...)
	return res1, nil
}

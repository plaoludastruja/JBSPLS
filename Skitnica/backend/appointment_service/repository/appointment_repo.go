package repository

import (
	"context"
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "appointmentsDB"
	COLLECTION = "appointments"
)

type AppointmentRepo struct {
	appointments *mongo.Collection
}

func NewAppointmentRepo(client *mongo.Client) domain.IAppointmentRepo {
	appointmentsCollection := client.Database(DATABASE).Collection(COLLECTION)
	return &AppointmentRepo{
		appointments: appointmentsCollection,
	}
}

func (store *AppointmentRepo) Get(id primitive.ObjectID) (*domain.Appointment, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AppointmentRepo) GetAll() ([]*domain.Appointment, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AppointmentRepo) Insert(appointment *domain.Appointment) error {
	result, err := store.appointments.InsertOne(context.TODO(), appointment)
	if err != nil {
		return err
	}
	appointment.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AppointmentRepo) Edit(appointment *domain.Appointment) error {
	filter := bson.M{"_id": appointment.Id}
	update := bson.M{"$set": bson.M{
		"accomodationId": appointment.AccomodationId,
		"start":          appointment.Start,
		"end":            appointment.End,
		"priceType":      appointment.PriceType,
		"price":          appointment.Price,
		"status":         appointment.Status,
	}}
	_, err := store.appointments.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *AppointmentRepo) DeleteAll() {
	store.appointments.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AppointmentRepo) Delete(id primitive.ObjectID) error {
	res, err := store.appointments.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return err
}

func (store *AppointmentRepo) filter(filter interface{}) ([]*domain.Appointment, error) {
	cursor, err := store.appointments.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AppointmentRepo) filterOne(filter interface{}) (appointment *domain.Appointment, err error) {
	result := store.appointments.FindOne(context.TODO(), filter)
	err = result.Decode(&appointment)
	return
}

func decode(cursor *mongo.Cursor) (appointments []*domain.Appointment, err error) {
	for cursor.Next(context.TODO()) {
		var appointment domain.Appointment
		err = cursor.Decode(&appointment)
		if err != nil {
			return
		}
		appointments = append(appointments, &appointment)
	}
	err = cursor.Err()
	return
}

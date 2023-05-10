package repository

import (
	"context"
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "accomodationsDB"
	COLLECTION = "accomodations"
)

type AccomodationRepo struct {
	accomodations *mongo.Collection
}

func NewAccomodationRepo(client *mongo.Client) domain.IAccomodationRepo {
	accomodationsCollection := client.Database(DATABASE).Collection(COLLECTION)
	return &AccomodationRepo{
		accomodations: accomodationsCollection,
	}
}

func (store *AccomodationRepo) Get(id primitive.ObjectID) (*domain.Accomodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccomodationRepo) GetAll() ([]*domain.Accomodation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccomodationRepo) Insert(accomodation *domain.Accomodation) error {
	result, err := store.accomodations.InsertOne(context.TODO(), accomodation)
	if err != nil {
		return err
	}
	accomodation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AccomodationRepo) Edit(accomodation *domain.Accomodation) error {
	filter := bson.M{"_id": accomodation.Id}
	update := bson.M{"$set": bson.M{
		"name":                accomodation.Name,
		"location":            accomodation.Location,
		"facilities":          accomodation.Facilities,
		"minNumberOfGuests":   accomodation.MinNumberOfGuests,
		"maxNumberOfGuests":   accomodation.MaxNumberOfGuests,
		"isAutomaticApproved": accomodation.IsAutomaticApproved,
	}}
	_, err := store.accomodations.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *AccomodationRepo) DeleteAll() {
	store.accomodations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AccomodationRepo) Delete(id primitive.ObjectID) error {
	res, err := store.accomodations.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return err
}

func (store *AccomodationRepo) filter(filter interface{}) ([]*domain.Accomodation, error) {
	cursor, err := store.accomodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccomodationRepo) filterOne(filter interface{}) (accomodation *domain.Accomodation, err error) {
	result := store.accomodations.FindOne(context.TODO(), filter)
	err = result.Decode(&accomodation)
	return
}

func decode(cursor *mongo.Cursor) (accomodations []*domain.Accomodation, err error) {
	for cursor.Next(context.TODO()) {
		var accomodation domain.Accomodation
		err = cursor.Decode(&accomodation)
		if err != nil {
			return
		}
		accomodations = append(accomodations, &accomodation)
	}
	err = cursor.Err()
	return
}

/*
func checkExisting(name string, where []domain.Accomodation) (idx int) {
	for _, v := range where {
		if v.Name == name {
			return 0
		}
	}
	return 1
}
*/

package repository

import (
	"context"
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "accomodationRatingsDB"
	COLLECTION = "accomodationRatings"
)

type AccomodationRatingRepo struct {
	accomodationRatings *mongo.Collection
}

func NewAccomodationRatingRepo(client *mongo.Client) domain.IAccomodationRatingRepo {
	accomodationRatingsCollection := client.Database(DATABASE).Collection(COLLECTION)
	return &AccomodationRatingRepo{
		accomodationRatings: accomodationRatingsCollection,
	}
}

func (store *AccomodationRatingRepo) Get(id primitive.ObjectID) (*domain.AccomodationRating, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccomodationRatingRepo) GetAll() ([]*domain.AccomodationRating, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccomodationRatingRepo) Insert(accomodationRating *domain.AccomodationRating) error {
	result, err := store.accomodationRatings.InsertOne(context.TODO(), accomodationRating)
	if err != nil {
		return err
	}
	accomodationRating.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AccomodationRatingRepo) Edit(accomodationRating *domain.AccomodationRating) error {
	filter := bson.M{"_id": accomodationRating.Id}
	update := bson.M{"$set": bson.M{
		"email":          accomodationRating.Email,
		"accomodationId": accomodationRating.AccomodationId,
		"rating":         accomodationRating.Rating,
	}}
	_, err := store.accomodationRatings.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *AccomodationRatingRepo) DeleteAll() {
	store.accomodationRatings.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AccomodationRatingRepo) Delete(id primitive.ObjectID) error {
	res, err := store.accomodationRatings.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return err
}

func (store *AccomodationRatingRepo) filter(filter interface{}) ([]*domain.AccomodationRating, error) {
	cursor, err := store.accomodationRatings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccomodationRatingRepo) filterOne(filter interface{}) (accomodationRating *domain.AccomodationRating, err error) {
	result := store.accomodationRatings.FindOne(context.TODO(), filter)
	err = result.Decode(&accomodationRating)
	return
}

func decode(cursor *mongo.Cursor) (accomodationRatings []*domain.AccomodationRating, err error) {
	for cursor.Next(context.TODO()) {
		var accomodationRating domain.AccomodationRating
		err = cursor.Decode(&accomodationRating)
		if err != nil {
			return
		}
		accomodationRatings = append(accomodationRatings, &accomodationRating)
	}
	err = cursor.Err()
	return
}

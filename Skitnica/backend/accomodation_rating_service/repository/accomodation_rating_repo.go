package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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
	driver              neo4j.DriverWithContext
}

func NewAccomodationRatingRepo(client *mongo.Client) domain.IAccomodationRatingRepo {
	accomodationRatingsCollection := client.Database(DATABASE).Collection(COLLECTION)
	uri := os.Getenv("NEO4J_DB")
	user := os.Getenv("NEO4J_USERNAME")
	pass := os.Getenv("NEO4J_PASS")
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		return nil
	}
	return &AccomodationRatingRepo{
		accomodationRatings: accomodationRatingsCollection,
		driver:              driver,
	}
}

func (pr *AccomodationRatingRepo) CheckConnection() {
	ctx := context.Background()
	err := pr.driver.VerifyConnectivity(ctx)
	if err != nil {
		return
	}
}

func (pr *AccomodationRatingRepo) CloseDriverConnection(ctx context.Context) {
	pr.driver.Close(ctx)
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
		"date":           accomodationRating.Date,
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

func (acr *AccomodationRatingRepo) WriteAccomodationRating(accomodationRating *domain.AccomodationRating) error {
	ctx := context.Background()
	session := acr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteWrite for write transactions (Create/Update/Delete)
	savedAccomodationRating, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (ar:AccomodationRating) SET ar.email = $email, ar.accomodationId = $accomodationId, ar.date = $date, ar.rating = $rating RETURN ar.email + ', from node ' + id(ar)",
				map[string]any{"email": accomodationRating.Email, "rating": accomodationRating.Rating, "accomodationId": accomodationRating.AccomodationId, "date": accomodationRating.Date})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		fmt.Println("Error inserting Accomodation Rating:", err)
		return err
	}
	fmt.Println(savedAccomodationRating.(string))
	return nil
}

package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/domain"
	accomodationProto "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	notificationProto "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DATABASE   = "accomodationRatingsDB"
	COLLECTION = "accomodationRatings"
)

type AccomodationRatingRepo struct {
	accomodationRatings *mongo.Collection
	neo4jSession        neo4j.Session
}

func NewAccomodationRatingRepo(client *mongo.Client, neo4jSession neo4j.Session) domain.IAccomodationRatingRepo {
	accomodationRatingsCollection := client.Database(DATABASE).Collection(COLLECTION)
	return &AccomodationRatingRepo{
		accomodationRatings: accomodationRatingsCollection,
		neo4jSession:        neo4jSession,
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
	fmt.Println("1")
	query := "CREATE (p:AccomodationRating {Email: $email, AccomodationId: $accomodationId, Rating: $rating, Date: $date })"
	fmt.Println("2")
	parameters := map[string]interface{}{
		"email":          accomodationRating.Email,
		"accomodationId": accomodationRating.AccomodationId,
		"rating":         strconv.Itoa(int(accomodationRating.Rating)),
		"date":           accomodationRating.Date,
	}
	fmt.Println("3")
	result1, err1 := store.neo4jSession.Run(query, parameters)
	if err1 != nil {
		fmt.Println("4b")
		panic(err1)
	}

	fmt.Println("4")
	// Check if the query executed successfully
	if result1.Err() != nil {
		fmt.Println("5b")
		panic(result1.Err())
	}

	fmt.Println("5")
	// ako se ocjeni hostov smjestaj, salje se notifikacija na notifiaction_service
	if err == nil {
		notificationEndpoint := fmt.Sprintf("%s:%s", "notification_service", "8000")
		conn, err := grpc.Dial(notificationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
		}
		notificationClient := notificationProto.NewNotificationServiceClient(conn)

		accomodationEndpoint := fmt.Sprintf("%s:%s", "accomodation_service", "8000")
		conn2, err2 := grpc.Dial(accomodationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err2 != nil {
			log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
		}
		accomodationClient := accomodationProto.NewAccomodationServiceClient(conn2)

		accomodation, err2 := accomodationClient.Get(context.TODO(), &accomodationProto.GetRequest{Id: accomodationRating.AccomodationId})

		notificationPb := notificationProto.Notification{
			Id:       primitive.NewObjectID().Hex(),
			Receiver: accomodation.Accomodation.HostUsername,
			Sender:   "SYSTEM",
			Subject:  "GRADE_ACCOMODATION",
			Message:  "Vaš smještaj je ocjenjen  " + " sa ocjenom " + strconv.Itoa(int(accomodationRating.Rating)),
			IsRead:   "false",
		}

		notification, err := notificationClient.CreateNotification(context.TODO(), &notificationProto.CreateNotificationRequest{Notification: &notificationPb})
		fmt.Println(notification.Notification)
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

func (store *AccomodationRatingRepo) GetRecommended(email string) ([]string, error) {
	query := `MATCH (r:AccomodationRating) 
	WHERE r.email = $email
	RETURN r.id as id, r.email as email, r.accomodationId as accomodationId, r.rating as rating, r.date as date`
	fmt.Println("2")
	fmt.Println("3")
	result1, err1 := store.neo4jSession.Run(query, map[string]any{"email": email})
	if err1 != nil {
		fmt.Println("4b")
		panic(err1)
	}

	var ratings []*domain.AccomodationRating
	for result1.Next() {
		record := result1.Record()
		id, _ := record.Get("id")
		email, _ := record.Get("email")
		accomodationId, _ := record.Get("accomodationId")
		rating, _ := record.Get("rating")
		date, _ := record.Get("date")
		ratings = append(ratings, &domain.AccomodationRating{
			Id:             id.(primitive.ObjectID),
			Email:          email.(string),
			AccomodationId: accomodationId.(string),
			Rating:         rating.(int32),
			Date:           date.(string),
		})
	}
	var ratings3 []*domain.AccomodationRating
	for _, rat := range ratings {
		query1 := `MATCH (r:AccomodationRating) 
		WHERE r.accomodationId = $accomodationId AND r.rating BETWEEN $n AND $m
		RETURN r.id as id, r.email as email, r.accomodationId as accomodationId, r.rating as rating, r.date as date`
		result2, err2 := store.neo4jSession.Run(query1, map[string]any{"accomodationId": rat.AccomodationId, "n": rat.Rating - 1, "m": rat.Rating + 1})
		if err2 != nil {
			panic(err2)
		}

		var ratings2 []*domain.AccomodationRating
		for result2.Next() {
			record := result2.Record()
			id, _ := record.Get("id")
			email, _ := record.Get("email")
			accomodationId, _ := record.Get("accomodationId")
			rating, _ := record.Get("rating")
			date, _ := record.Get("date")
			ratings2 = append(ratings2, &domain.AccomodationRating{
				Id:             id.(primitive.ObjectID),
				Email:          email.(string),
				AccomodationId: accomodationId.(string),
				Rating:         rating.(int32),
				Date:           date.(string),
			})
		}
		ratings3 = append(ratings3, ratings2...)

	}
	var accomodations []string
	for _, rat := range ratings3 {
		if rat.Rating == 4 || rat.Rating == 5 {
			accomodations = append(accomodations, rat.AccomodationId)
		}

	}

	return accomodations, nil
}

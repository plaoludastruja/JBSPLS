package repository

import (
	"context"
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/hostmark_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "hostMarksDB"
	COLLECTION = "hostMarks"
)

type HostMarkRepo struct {
	hostMarks *mongo.Collection
}

func NewHostMarkRepo(client *mongo.Client) domain.IHostMarkRepo {
	hostMarksCollection := client.Database(DATABASE).Collection(COLLECTION)
	return &HostMarkRepo{
		hostMarks: hostMarksCollection,
	}
}

func (store *HostMarkRepo) Get(id primitive.ObjectID) (*domain.HostMark, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *HostMarkRepo) GetAll() ([]*domain.HostMark, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *HostMarkRepo) Insert(user *domain.HostMark) error {
	result, err := store.hostMarks.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *HostMarkRepo) Edit(hostMark *domain.HostMark) error {
	filter := bson.M{"_id": hostMark.Id}
	update := bson.M{"$set": bson.M{
		"username":     hostMark.Username,
		"grade":        hostMark.Grade,
		"hostUsername": hostMark.HostUsername,
		"dateTime"	: hostMark.DateTime,
	}}
	_, err := store.hostMarks.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *HostMarkRepo) DeleteAll() {
	store.hostMarks.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *HostMarkRepo) Delete(id primitive.ObjectID) error {
	res, err := store.hostMarks.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return err
}

func (store *HostMarkRepo) filter(filter interface{}) ([]*domain.HostMark, error) {
	cursor, err := store.hostMarks.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *HostMarkRepo) filterOne(filter interface{}) (user *domain.HostMark, err error) {
	result := store.hostMarks.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func decode(cursor *mongo.Cursor) (hostMarks []*domain.HostMark, err error) {
	for cursor.Next(context.TODO()) {
		var hostMark domain.HostMark
		err = cursor.Decode(&hostMark)
		if err != nil {
			return
		}
		hostMarks = append(hostMarks, &hostMark)
	}
	err = cursor.Err()
	return
}

func (store *HostMarkRepo) GetByUsername(username string, hostUsername string) ([]*domain.HostMark, error) {
	filter := bson.M{"username": username, "hostUsername": hostUsername}
	return store.filter(filter)
}

func (store *HostMarkRepo) GetAllByHostUsername(hostUsername string) ([]*domain.HostMark, error) {
	filter := bson.M{"hostUsername": hostUsername}
	return store.filter(filter)
}

func (store *HostMarkRepo) GetByHost(hostUsername string) ([]*domain.HostMark, error) {
	filter := bson.M{"hostUsername": hostUsername}
	return store.filter(filter)
}
package repository

import (
	"context"
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "notificationDB"
	COLLECTION  = "notifications"
	COLLECTION2 = "notificationFilters"
)

type NotificationRepo struct {
	notifications       *mongo.Collection
	notificationFilters *mongo.Collection
}

func NewNotificationRepo(client *mongo.Client) domain.INotificationRepo {
	notificationsCollection := client.Database(DATABASE).Collection(COLLECTION)
	notificationFiltersCollection := client.Database(DATABASE).Collection(COLLECTION2)
	return &NotificationRepo{
		notifications:       notificationsCollection,
		notificationFilters: notificationFiltersCollection,
	}
}

func (store *NotificationRepo) Get(id primitive.ObjectID) (*domain.Notification, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *NotificationRepo) GetBySender(username string) ([]*domain.Notification, error) {
	filter := bson.M{"sender": username}
	return store.filter(filter)
}

func (store *NotificationRepo) GetByReceiver(username string) ([]*domain.Notification, error) {
	filter := bson.M{"receiver": username}
	return store.filter(filter)
}

func (store *NotificationRepo) GetAll() ([]*domain.Notification, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *NotificationRepo) Insert(notification *domain.Notification) error {
	result, err := store.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *NotificationRepo) Edit(notification *domain.Notification) error {
	filter := bson.M{"_id": notification.Id}
	update := bson.M{"$set": bson.M{
		"isRead": "true",
	}}
	_, err := store.notifications.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *NotificationRepo) DeleteAll() {
	store.notifications.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *NotificationRepo) Delete(id primitive.ObjectID) error {
	res, err := store.notifications.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return err
}

func (store *NotificationRepo) filter(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notifications.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *NotificationRepo) filterOne(filter interface{}) (notification *domain.Notification, err error) {
	result := store.notifications.FindOne(context.TODO(), filter)
	err = result.Decode(&notification)
	return
}

func decode(cursor *mongo.Cursor) (notifications []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var notification domain.Notification
		err = cursor.Decode(&notification)
		if err != nil {
			return
		}
		notifications = append(notifications, &notification)
	}
	err = cursor.Err()
	return
}

func (store *NotificationRepo) InsertNotificationFilters(notificationFilter *domain.NotificationFilter) error {
	result, err := store.notificationFilters.InsertOne(context.TODO(), notificationFilter)
	if err != nil {
		return err
	}
	notificationFilter.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *NotificationRepo) EditNotificationFilter(notificationFilter *domain.NotificationFilter) error {
	filter := bson.M{"_id": notificationFilter.Id}
	update := bson.M{"$set": bson.M{
		"reservation": notificationFilter.Reservation,
		"rating":      notificationFilter.Rating,
		"super":       notificationFilter.Super,
	}}
	_, err := store.notificationFilters.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *NotificationRepo) GetNotificationFilterByUsername(username string) (*domain.NotificationFilter, error) {
	filter := bson.M{"username": username}
	return store.filterOne1(filter)
}

func (store *NotificationRepo) filterOne1(filter interface{}) (notificationFilter *domain.NotificationFilter, err error) {
	result := store.notificationFilters.FindOne(context.TODO(), filter)
	err = result.Decode(&notificationFilter)
	return
}

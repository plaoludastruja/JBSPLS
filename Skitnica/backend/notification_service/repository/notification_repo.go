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
	DATABASE   = "notificationDB"
	COLLECTION = "notifications"
)

type NotificationRepo struct {
	notifications *mongo.Collection
}

func NewNotificationRepo(client *mongo.Client) domain.INotificationRepo {
	notificationsCollection := client.Database(DATABASE).Collection(COLLECTION)
	return &NotificationRepo{
		notifications: notificationsCollection,
	}
}

func (store *NotificationRepo) Get(id primitive.ObjectID) (*domain.Notification, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *NotificationRepo) GetBySender(username string) (*domain.Notification, error) {
	filter := bson.M{"sender": username}
	return store.filterOne(filter)
}

func (store *NotificationRepo) GetByReceiver(username string) (*domain.Notification, error) {
	filter := bson.M{"receiver": username}
	return store.filterOne(filter)
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
		"isRead": notification.IsRead,
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

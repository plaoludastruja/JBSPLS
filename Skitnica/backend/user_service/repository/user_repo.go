package repository

import (
	"context"
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "usersDB"
	COLLECTION = "users"
)

type UserRepo struct {
	users *mongo.Collection
}

func NewUserRepo(client *mongo.Client) domain.IUserRepo {
	usersCollection := client.Database(DATABASE).Collection(COLLECTION)
	return &UserRepo{
		users: usersCollection,
	}
}

func (store *UserRepo) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserRepo) GetByUsername(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserRepo) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserRepo) Insert(user *domain.User) error {
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserRepo) Edit(user *domain.User) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{
		"username":   user.Username,
		"password":   user.Password,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"address":    user.Address,
		"apiKey":     user.ApiKey,
	}}
	_, err := store.users.UpdateOne(context.TODO(), filter, update)

	return err
}

func (store *UserRepo) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserRepo) Delete(id primitive.ObjectID) error {
	res, err := store.users.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return err
}

func (store *UserRepo) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserRepo) filterOne(filter interface{}) (user *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var user domain.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}

func (store *UserRepo) GetHosts() ([]*domain.User, error) {
	filter := bson.M{"role": "HOST"}
	return store.filter(filter)
}

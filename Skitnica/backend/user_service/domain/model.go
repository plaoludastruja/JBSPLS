package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Color struct {
	Code string `bson:"code"`
	Name string `bson:"name"`
}

type Product struct {
	Id            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	ClothingBrand string             `bson:"clothing_brand"`
	Colors        []Color            `bson:"colors"`
}

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
	FirstName string             `json:"firstName" bson:"first_name" binding:"required"`
	LastName  string             `json:"lastName" bson:"last_name" binding:"required"`
	Role      string             `json:"role" bson:"role" binding:"required"`
	Address   string             `json:"address" bson:"address" binding:"required"`
}

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
	Id   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name" binding:"required"`
}

package DTO

type SearchDTO struct {
	StartPlace string `json:"startPlace" bson:"start_place" binding:"required"`
	EndPlace   string `json:"endPlace" bson:"end_place" binding:"required"`
}

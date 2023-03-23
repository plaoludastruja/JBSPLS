package DTO

type SearchDTO struct {
	StartPlace string `json:"startPlace" bson:"start_place"`
	EndPlace   string `json:"endPlace" bson:"end_place"`
}

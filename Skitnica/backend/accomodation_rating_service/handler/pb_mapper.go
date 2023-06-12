package handler

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/domain"
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_rating_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAccomodationRating(accomodationRating *domain.AccomodationRating) *pb.AccomodationRating {
	accomodationRatingPb := &pb.AccomodationRating{
		Id:             accomodationRating.Id.Hex(),
		Email:          accomodationRating.Email,
		AccomodationId: accomodationRating.AccomodationId,
		Rating:         accomodationRating.Rating,
		Date:           accomodationRating.Date,
	}
	return accomodationRatingPb
}

func mapAccomodationRatingPb(accomodationRatingPb *pb.AccomodationRating) *domain.AccomodationRating {
	accomodationRatingPbId, _ := primitive.ObjectIDFromHex(accomodationRatingPb.Id)
	accomodationRating := &domain.AccomodationRating{
		Id:             accomodationRatingPbId,
		Email:          accomodationRatingPb.Email,
		AccomodationId: accomodationRatingPb.AccomodationId,
		Rating:         accomodationRatingPb.Rating,
		Date:           accomodationRatingPb.Date,
	}
	return accomodationRating
}

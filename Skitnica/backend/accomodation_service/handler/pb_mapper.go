package handler

import (
	"strconv"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/domain"
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAccomodation(accomodation *domain.Accomodation) *pb.Accomodation {
	accomodationPb := &pb.Accomodation{
		Id:                  accomodation.Id.Hex(),
		Name:                accomodation.Name,
		Location:            accomodation.Location,
		Facilities:          accomodation.Facilities,
		MinNumberOfGuests:   accomodation.MinNumberOfGuests,
		MaxNumberOfGuests:   accomodation.MaxNumberOfGuests,
		IsAutomaticApproved: strconv.FormatBool(accomodation.IsAutomaticApproved),
		HostUsername:        accomodation.HostUsername,
		Image:               []byte(accomodation.Image),
	}
	return accomodationPb
}

func mapAccomodationPb(accomodationPb *pb.Accomodation) *domain.Accomodation {
	accomodationPbId, _ := primitive.ObjectIDFromHex(accomodationPb.Id)
	isAutomaticApprovedBool, _ := strconv.ParseBool(accomodationPb.IsAutomaticApproved)
	accomodation := &domain.Accomodation{
		Id:                  accomodationPbId,
		Name:                accomodationPb.Name,
		Location:            accomodationPb.Location,
		Facilities:          accomodationPb.Facilities,
		MinNumberOfGuests:   accomodationPb.MinNumberOfGuests,
		MaxNumberOfGuests:   accomodationPb.MaxNumberOfGuests,
		IsAutomaticApproved: isAutomaticApprovedBool,
		HostUsername:        accomodationPb.HostUsername,
		Image:               accomodationPb.Image,
	}
	return accomodation
}

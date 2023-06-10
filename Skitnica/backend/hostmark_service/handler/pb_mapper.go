package handler

import (
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/hostmark_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/hostmark_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapHostMark(hostMark *domain.HostMark) *pb.HostMark {
	hostMarkPb := &pb.HostMark{
		Id:           hostMark.Id.Hex(),
		Username:     hostMark.Username,
		Grade:        hostMark.Grade,
		HostUsername: hostMark.HostUsername,
	}
	return hostMarkPb
}

func mapHostMarkPb(hostMarkPb *pb.HostMark) *domain.HostMark {
	hostMarkPbId, _ := primitive.ObjectIDFromHex(hostMarkPb.Id)
	hostMark := &domain.HostMark{
		Id:           hostMarkPbId,
		Username:     hostMarkPb.Username,
		HostUsername: hostMarkPb.HostUsername,
	}
	return hostMark
}

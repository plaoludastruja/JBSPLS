package handler

import (
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/user_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       user.Id.Hex(),
		Username: user.Name,
	}
	return userPb
}

func mapUserPb(userPb *pb.User) *domain.User {
	userPbId, _ := primitive.ObjectIDFromHex(userPb.Id)
	user := &domain.User{
		Id:   userPbId,
		Name: userPb.Username,
	}
	return user
}

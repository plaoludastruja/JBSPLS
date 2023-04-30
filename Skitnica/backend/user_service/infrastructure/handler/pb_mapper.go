package handler

import (
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/user_service"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:   user.Id.Hex(),
		Name: user.Name,
	}
	return userPb
}

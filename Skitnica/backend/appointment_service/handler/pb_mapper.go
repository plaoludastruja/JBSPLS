package handler

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/domain"
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/appointment_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAppointment(appointment *domain.Appointment) *pb.Appointment {
	appointmentPb := &pb.Appointment{
		Id:             appointment.Id.Hex(),
		AccomodationId: appointment.AccomodationId.Hex(),
		Start:          appointment.Start,
		End:            appointment.End,
		PriceType:      appointment.PriceType,
		Price:          appointment.Price,
		Status:         appointment.Status,
	}
	return appointmentPb
}

func mapAppointmentPb(appointmentPb *pb.Appointment) *domain.Appointment {
	appointmentPbId, _ := primitive.ObjectIDFromHex(appointmentPb.Id)
	appointment := &domain.Appointment{
		Id:             appointmentPbId,
		AccomodationId: appointmentPb.AccomodationId,
		Start:          appointmentPb.Start,
		End:            appointmentPb.End,
		PriceType:      appointmentPb.PriceType,
		Price:          appointmentPb.Price,
		Status:         appointmentPb.Status,
	}
	return appointment
}

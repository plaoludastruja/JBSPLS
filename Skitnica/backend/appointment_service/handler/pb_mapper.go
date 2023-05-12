package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/domain"
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/appointment_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAppointment(appointment *domain.Appointment) *pb.Appointment {
	appointmentPb := &pb.Appointment{
		Id:             appointment.Id.Hex(),
		AccomodationId: appointment.AccomodationId,
		Start:          appointment.Start.String(),
		End:            appointment.End.String(),
		PriceType:      appointment.PriceType,
		Price:          appointment.Price,
		Status:         appointment.Status,
	}
	return appointmentPb
}

func mapAppointmentPb(appointmentPb *pb.Appointment) *domain.Appointment {
	appointmentPbId, _ := primitive.ObjectIDFromHex(appointmentPb.Id)
	const layout = "2006-01-02"
	fmt.Println(appointmentPb.Start)
	start, _ := time.Parse(layout, appointmentPb.Start)
	fmt.Println(start)
	end, _ := time.Parse(layout, appointmentPb.End) //"YYYY-MM-DD"
	appointment := &domain.Appointment{

		Id:             appointmentPbId,
		AccomodationId: appointmentPb.AccomodationId,
		Start:          start,
		End:            end,
		PriceType:      appointmentPb.PriceType,
		Price:          appointmentPb.Price,
		Status:         appointmentPb.Status,
	}
	if strings.Contains(appointmentPb.Start, "UTC") {
		fmt.Println("udje ovde")
		layout := "2006-01-02 15:04:05 -0700 MST"
		start, _ := time.Parse(layout, appointmentPb.Start)
		end, _ := time.Parse(layout, appointmentPb.End)
		appointment := &domain.Appointment{

			Id:             appointmentPbId,
			AccomodationId: appointmentPb.AccomodationId,
			Start:          start,
			End:            end,
			PriceType:      appointmentPb.PriceType,
			Price:          appointmentPb.Price,
			Status:         appointmentPb.Status,
		}
		return appointment
	}
	return appointment
}

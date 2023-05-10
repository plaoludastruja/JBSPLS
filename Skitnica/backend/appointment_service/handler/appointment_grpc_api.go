package handler

import (
	"context"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/service"
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/appointment_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppointmentHandler struct {
	pb.UnimplementedAppointmentServiceServer
	service *service.AppointmentService
}

func NewAppointmentHandler(service *service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		service: service,
	}
}

func (handler *AppointmentHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	appointment, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	appointmentPb := mapAppointment(appointment)
	response := &pb.GetResponse{
		Appointment: appointmentPb,
	}
	return response, nil
}

func (handler *AppointmentHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	appointments, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Appointments: []*pb.Appointment{},
	}
	for _, appointment := range appointments {
		current := mapAppointment(appointment)
		response.Appointments = append(response.Appointments, current)
	}
	return response, nil
}

func (handler *AppointmentHandler) CreateAppointment(ctx context.Context, request *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error) {
	appointment := mapAppointmentPb(request.Appointment)
	err := handler.service.Insert(*appointment)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppointmentResponse{
		Appointment: mapAppointment(appointment),
	}, nil
}

func (handler *AppointmentHandler) EditAppointment(ctx context.Context, request *pb.EditAppointmentRequest) (*pb.EditAppointmentResponse, error) {
	appointment := mapAppointmentPb(request.Appointment)
	err := handler.service.Edit(*appointment)
	if err != nil {
		return nil, err
	}
	return &pb.EditAppointmentResponse{
		Appointment: mapAppointment(appointment),
	}, nil
}

func (handler *AppointmentHandler) DeleteAppointment(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	errr := handler.service.Delete(objectId)
	if errr != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (handler *AppointmentHandler) GetByAccomodation(ctx context.Context, request *pb.GetAccomodationRequest) (*pb.GetAccomodationResponse, error) {
	appointments, err := handler.service.GetByAccomodation(request.AccomodationId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAccomodationResponse{
		Appointments: []*pb.Appointment{},
	}
	for _, appointment := range appointments {
		current := mapAppointment(appointment)
		response.Appointments = append(response.Appointments, current)
	}
	return response, nil
}

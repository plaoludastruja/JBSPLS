package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/infrastructure/service"
	hostmark "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/hostmark_service/generated"
	notificationProto "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	reservation "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BestHosthHandler struct {
	reservationClientAddress string
	hostmarkClientAddress    string
}

func NewBestHosthHandler(reservationClientAddress, hostmarkClientAddress string) Handler {
	return &BestHosthHandler{
		reservationClientAddress: reservationClientAddress,
		hostmarkClientAddress:    hostmarkClientAddress,
	}
}

func (handler *BestHosthHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/isBestHost/{hostUsername}", handler.IsHostBestHost)
	if err != nil {
		panic(err)
	}
}

func (handler *BestHosthHandler) IsHostBestHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	hostUsername := pathParams["hostUsername"]
	fmt.Println("hostUsername", hostUsername)

	// ako se ocjeni hostov smjestaj, salje se notifikacija na notifiaction_service

	notificationEndpoint := fmt.Sprintf("%s:%s", "notification_service", "8000")
	conn, err := grpc.Dial(notificationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	notificationClient := notificationProto.NewNotificationServiceClient(conn)

	reservationClient := service.NewReservationClient(handler.reservationClientAddress)
	reservationsResult, _ := reservationClient.IsHostBestHost(context.TODO(), &reservation.IsHostBestHostRequest{HostUsername: hostUsername})
	if reservationsResult.IsBestHost == false {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("false"))
		notificationPb := notificationProto.Notification{
			Id:       primitive.NewObjectID().Hex(),
			Receiver: hostUsername,
			Sender:   "SYSTEM",
			Subject:  "SUPERHOST",
			Message:  " Nemate status superhosta. ",
			IsRead:   "false",
		}

		notification, _ := notificationClient.CreateNotification(context.TODO(), &notificationProto.CreateNotificationRequest{Notification: &notificationPb})
		fmt.Println(notification.Notification)
		return
	}

	fmt.Println(reservationsResult.IsBestHost)

	hostmarkClient := service.NewHostMarkClient(handler.hostmarkClientAddress)
	hostmarkResult, _ := hostmarkClient.IsHostBestHost(context.TODO(), &hostmark.IsHostBestHostRequest{HostUsername: hostUsername})
	if hostmarkResult.IsBestHost == false {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("false"))
		notificationPb := notificationProto.Notification{
			Id:       primitive.NewObjectID().Hex(),
			Receiver: hostUsername,
			Sender:   "SYSTEM",
			Subject:  "SUPERHOST",
			Message:  " Nemate status superhosta. ",
			IsRead:   "false",
		}

		notification, _ := notificationClient.CreateNotification(context.TODO(), &notificationProto.CreateNotificationRequest{Notification: &notificationPb})
		fmt.Println(notification.Notification)
		return
	}

	notificationPb := notificationProto.Notification{
		Id:       primitive.NewObjectID().Hex(),
		Receiver: hostUsername,
		Sender:   "SYSTEM",
		Subject:  "SUPERHOST",
		Message:  " Imate status superhosta. ",
		IsRead:   "false",
	}

	notification, _ := notificationClient.CreateNotification(context.TODO(), &notificationProto.CreateNotificationRequest{Notification: &notificationPb})
	fmt.Println(notification.Notification)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("true"))

}

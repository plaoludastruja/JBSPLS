package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/infrastructure/service"
	hostmark "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/hostmark_service/generated"
	reservation "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
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

	/*if location == "" || guestNumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
		}
	num, err := strconv.Atoi(guestNumber)
	if err != nil {
		fmt.Println("Greška pri parsiranju:", err)
		return
	}
	num2 := int32(num)
	fmt.Println(num)
	//searchResult := &domain.SearchResult{Name: ""}
	accomodationClient := service.NewAccomodationClient(handler.accomodationClientAddress)
	accomodations, err := accomodationClient.Search(context.TODO(), &accomodation.SearchRequest{Location: location, GuestNumber: num2})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}*/

	reservationClient := service.NewReservationClient(handler.reservationClientAddress)
	reservationsResult, _ := reservationClient.IsHostBestHost(context.TODO(), &reservation.IsHostBestHostRequest{HostUsername: hostUsername})
	if reservationsResult.IsBestHost == false {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("false"))
		return
	}

	fmt.Println(reservationsResult.IsBestHost)

	hostmarkClient := service.NewHostMarkClient(handler.hostmarkClientAddress)
	hostmarkResult, _ := hostmarkClient.IsHostBestHost(context.TODO(), &hostmark.IsHostBestHostRequest{HostUsername: hostUsername})
	if hostmarkResult.IsBestHost == false {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("false"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("true"))

}

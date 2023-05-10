package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/infrastructure/service"
	accomodation "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
)

type SearchHandler struct {
	accomodationClientAddress string
	appointmentClientAddress  string
	reservationClientAddress  string
}

func NewSearchHandler(accomodationClientAddress, appointmentClientAddress, reservationClientAddress string) Handler {
	return &SearchHandler{
		accomodationClientAddress: accomodationClientAddress,
		appointmentClientAddress:  appointmentClientAddress,
		reservationClientAddress:  reservationClientAddress,
	}
}

func (handler *SearchHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/accomodation/search/{location}/{guestNumber}", handler.Search)
	if err != nil {
		panic(err)
	}
}

func (handler *SearchHandler) Search(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	location := pathParams["location"]
	guestNumber := pathParams["guestNumber"]
	if location == "" || guestNumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	num, err := strconv.Atoi(guestNumber)
	if err != nil {
		fmt.Println("Greška pri parsiranju:", err)
		return
	}
	fmt.Println(num)
	//searchResult := &domain.SearchResult{Name: ""}
	accomodationClient := service.NewAccomodationClient(handler.accomodationClientAddress)
	accomodations, err := accomodationClient.GetAll(context.TODO(), &accomodation.GetAllRequest{})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, err := json.Marshal(accomodations)
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

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
	reservation "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
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
	err := mux.HandlePath("GET", "/accomodation/search/{location}/{guestNumber}/{startDay}/{startMonth}/{startYear}/{endDay}/{endMonth}/{endYear}", handler.Search)
	if err != nil {
		panic(err)
	}
}

func (handler *SearchHandler) Search(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	location := pathParams["location"]
	guestNumber := pathParams["guestNumber"]

	startDayString := pathParams["startDay"]
	fmt.Println("startDayString", startDayString)

	startMonthString := pathParams["startMonth"]
	fmt.Println("startMonthString", startMonthString)

	startYearString := pathParams["startYear"]
	fmt.Println("startYearString", startYearString)

	endDayString := pathParams["endDay"]
	fmt.Println("endDayString", endDayString)

	endMonthString := pathParams["endMonth"]
	fmt.Println("endMonthString", endMonthString)

	endYearString := pathParams["endYear"]
	fmt.Println("endYearString", endYearString)

	if location == "" || guestNumber == "" {
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
	}

	reservationClient := service.NewReservationClient(handler.reservationClientAddress)
	reservations, _ := reservationClient.Search(context.TODO(), &reservation.SearchRequest{StartDay: startDayString, StartMonth: startMonthString, StartYear: startYearString, EndDay: endDayString, EndMonth: endMonthString, EndYear: endYearString})

	fmt.Println(reservations.Reservations)

	find := false

	result := []accomodation.Accomodation{}
	for _, accomodation := range accomodations.Accomodations {
		for _, reservation := range reservations.Reservations {
			fmt.Println("Usao u reservations for")
			if accomodation.Id == reservation.AccomodationId {
				fmt.Println("Usao u reservations if")
				find = true
				fmt.Println("Find:", find)
			}
		}
		if !find {
			fmt.Println("Usao u append if")
			result = append(result, *accomodation)
			fmt.Println(result)
		}
		find = false
		fmt.Println("Find:", find)
	}

	/*appointmentClient := service.NewAppointmentClient(handler.appointmentClientAddress)
	appointments, _ := appointmentClient.Search(context.TODO(), &reservation.SearchRequest{StartDay: startDayString, StartMonth: startMonthString, StartYear: startYearString, EndDay: endDayString, EndMonth: endMonthString, EndYear: endYearString})*/

	response, err := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

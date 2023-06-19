package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/domain"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/infrastructure/service"
	accomodation "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	appointment "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/appointment_service/generated"
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
	//result := make([]*domain.Accomodation, 0)
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

	appointmentClient := service.NewAppointmentClient(handler.appointmentClientAddress)
	appointments, _ := appointmentClient.GetAll(context.TODO(), &appointment.GetAllRequest{})

	const layout = "2006-01-02 15:04:05 -0700 MST"

	startYearInt, _ := strconv.Atoi(startYearString)
	startMonthInt, _ := strconv.Atoi(startMonthString)
	startDayInt, _ := strconv.Atoi(startDayString)

	endYearInt, _ := strconv.Atoi(endYearString)
	endMonthInt, _ := strconv.Atoi(endMonthString)
	endDayInt, _ := strconv.Atoi(endDayString)

	start := time.Date(startYearInt, time.Month(startMonthInt), startDayInt, 00, 00, 00, 999999999, time.UTC)
	end := time.Date(endYearInt, time.Month(endMonthInt), endDayInt, 00, 00, 00, 999999999, time.UTC)

	duration := end.Sub(start)
	days := int32(duration.Hours() / 24)

	result2 := []accomodation.Accomodation{}
	res := []domain.SearchResult{}

	fmt.Println("pre fora sa app ", result)
	for _, accomodation := range result {
		for _, appointment := range appointments.Appointments {
			if appointment.AccomodationId == accomodation.Id {
				fmt.Println("Usao u appointment for")
				appointmentStart, _ := time.Parse(layout, appointment.Start)
				appointmentEnd, _ := time.Parse(layout, appointment.End)
				fmt.Println("appointmentStart", appointmentStart)
				fmt.Println("appointmentEnd", appointmentEnd)
				if appointmentStart.Before(start) && appointmentEnd.After(end) {
					fmt.Println("Usao u appointment if")
					result2 = append(result2, accomodation)

					fmt.Println("days", days)

					gn := int32(1)
					if appointment.PriceType == "PerPerson" {
						gn = num2
					}

					prices := []int32{appointment.Price}
					priceTypes := []string{appointment.PriceType}

					sr := domain.SearchResult{AccomodationId: accomodation.Id,
						Name:                accomodation.Name,
						Location:            accomodation.Location,
						Facilities:          accomodation.Facilities,
						MinNumberOfGuests:   accomodation.MinNumberOfGuests,
						MaxNumberOfGuests:   accomodation.MaxNumberOfGuests,
						TotalPrice:          appointment.Price * days * gn,
						IsAutomaticApproved: accomodation.IsAutomaticApproved,
						Prices:              prices,
						PriceType:           priceTypes,
						HostUsername:        accomodation.HostUsername,
						Image:               accomodation.Image}
					found := false
					for _, searchResult := range res {
						if searchResult.AccomodationId == sr.AccomodationId {
							found = true
						}
					}
					if !found {
						res = append(res, sr)
					}
					fmt.Println(result2)
				} else if appointmentStart.Before(start) && appointmentEnd.Before(end) && appointmentEnd.After(start) {
					fmt.Println("Usao u appointment else if 1")
					for _, app := range appointments.Appointments {
						if app.AccomodationId == accomodation.Id {
							appStart, _ := time.Parse(layout, app.Start)
							appEnd, _ := time.Parse(layout, app.End)
							if appointmentEnd.Equal(appStart) && appEnd.After(end) {
								result2 = append(result2, accomodation)
								duration1 := appointmentEnd.Sub(start)
								days1 := int32(duration1.Hours() / 24)
								fmt.Println("days1", days1)

								duration2 := end.Sub(appStart)
								days2 := int32(duration2.Hours() / 24)
								fmt.Println("days2", days2)
								fmt.Println("appointment.Price", appointment.Price)
								fmt.Println("app.Price", app.Price)

								gn1 := int32(1)
								gn2 := int32(1)
								if appointment.PriceType == "PerPerson" {
									gn1 = num2
								}

								if app.PriceType == "PerPerson" {
									gn2 = num2
								}

								prices := []int32{appointment.Price, app.Price}
								priceTypes := []string{appointment.PriceType, app.PriceType}
								sr := domain.SearchResult{AccomodationId: accomodation.Id,
									Name:                accomodation.Name,
									Location:            accomodation.Location,
									Facilities:          accomodation.Facilities,
									MinNumberOfGuests:   accomodation.MinNumberOfGuests,
									MaxNumberOfGuests:   accomodation.MaxNumberOfGuests,
									TotalPrice:          appointment.Price*(days1+1)*gn1 + app.Price*days2*gn2,
									IsAutomaticApproved: accomodation.IsAutomaticApproved,
									Prices:              prices,
									PriceType:           priceTypes,
									HostUsername:        accomodation.HostUsername,
									Image:               accomodation.Image}
								found := false
								for _, searchResult := range res {
									if searchResult.AccomodationId == sr.AccomodationId {
										found = true
									}
								}
								if !found {
									res = append(res, sr)
								}
							}
						}
					}
				} else if appointmentStart.After(start) && appointmentStart.Before(end) && appointmentEnd.After(end) {
					fmt.Println("Usao u appointment else if 2")
					for _, app2 := range appointments.Appointments {
						if appointment.AccomodationId == accomodation.Id {
							app2Start, _ := time.Parse(layout, app2.Start)
							app2End, _ := time.Parse(layout, app2.End)
							if appointmentStart.Equal(app2End) && app2Start.Before(appointmentStart) {
								result2 = append(result2, accomodation)

								duration1 := end.Sub(appointmentStart)
								days1 := int32(duration1.Hours() / 24)

								duration2 := app2End.Sub(start)
								days2 := int32(duration2.Hours() / 24)

								gn1 := int32(1)
								gn2 := int32(1)
								if appointment.PriceType == "PerPerson" {
									gn1 = num2
								}

								if app2.PriceType == "PerPerson" {
									gn2 = num2
								}

								prices := []int32{appointment.Price, app2.Price}
								priceTypes := []string{appointment.PriceType, app2.PriceType}
								sr := domain.SearchResult{AccomodationId: accomodation.Id,
									Name:                accomodation.Name,
									Location:            accomodation.Location,
									Facilities:          accomodation.Facilities,
									MinNumberOfGuests:   accomodation.MinNumberOfGuests,
									MaxNumberOfGuests:   accomodation.MaxNumberOfGuests,
									TotalPrice:          appointment.Price*days1*gn1 + app2.Price*days2*gn2,
									IsAutomaticApproved: accomodation.IsAutomaticApproved,
									Prices:              prices,
									PriceType:           priceTypes,
									HostUsername:        accomodation.HostUsername,
									Image:               accomodation.Image}
								found := false
								for _, searchResult := range res {
									if searchResult.AccomodationId == sr.AccomodationId {
										found = true
									}
								}
								if !found {
									res = append(res, sr)
								}
							}
						}
					}
				}
			}
		}
	}

	response, err := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

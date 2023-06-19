import { useEffect, useState } from "react";
import "./Reservations.css";
import reservationService from "../../services/reservation.service";
import Reservation from "../../model/Reservation";
import decodeToken from "../../services/auth.service";
import { MDBBtn, MDBCol, MDBContainer, MDBInput, MDBRow } from "mdb-react-ui-kit";
import StarRatings from "react-star-ratings";
import accomodationRatingService from "../../services/accomodationRating.service";
import AccomodationRating from "../../model/AccomodationRating";
import RecomendFlightsParam from "../../model/RecomendFlightsParam";
import flightService from "../../services/flights.service";
import { Flight } from "../../model/Flight";
import accomodationService from "../../services/accomodation.service";
import Accomodation from "../../model/Accomodation";
import userService from "../../services/user.service";
import ticketService from "../../services/ticket.service";

function Reservations() {
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [buttonClicked, setButtonClicked] = useState(false);
  const [oldRating, setOldRating] = useState<AccomodationRating | null>();
  const [rating, setRating] = useState(0);
  const [ratedAccomodation, setRatedAccomodation] = useState("");
  const [showRating, setShowRating] = useState(false);
  const [showRecommendForm, setshowRecommendForm] = useState(false);
  const [recommendParams, setRecommendParams] = useState<RecomendFlightsParam>({Departure:"", Arrival: ""});
  const [selectedReservation, setSelectedReservation] = useState<Reservation>({id:"",
                                                                              accomodationId:"",
                                                                              username: "",
                                                                              startDate:"",
                                                                              endDate: "",
                                                                              guestNumber: 0,
                                                                              status:"",
                                                                              hostUsername:""});
  const [flights1, setFlights1] = useState<Flight[]>([]);
  const [flights2, setFlights2] = useState<Flight[]>([]);
  const [accomodationFromReservation, setAccomodationFromReservation] = useState<Accomodation>({id: "",
  name: "",
  location: "",
  facilities: "",
  minNumberOfGuests: 0,
  maxNumberOfGuests: 0,
  isAutomaticApproved: "",
  hostUsername: "",
  image: ""});
  const [showRecommendTables, setShowRecommendTables] = useState(false);

  useEffect(() => {
    reservationService
      .getReservations(decodeToken()?.username)
      .then((response) => {
        setReservations(response.data.reservations);
      });
  }, [buttonClicked]);

  const handleOnClick = (reservation: Reservation) => {
    if (reservation.status === "APPROVED") {
      reservation.startDate = reservation.startDate.split(" ", 1)[0];
      reservation.endDate = reservation.endDate.split(" ", 1)[0];
      const reservationEndDate = new Date(reservation.endDate.split(" ", 1)[0]);
      const currentDate = new Date();
      if (reservationEndDate.getTime() < currentDate.getTime()) {
        setShowRating(true);
        setRatedAccomodation(reservation.accomodationId);
        accomodationRatingService
          .getAccomodationRatingsByAccomodationAndUser(
            reservation.accomodationId,
            decodeToken()?.username
          )
          .then((res) => {
            if (res.data.accomodationRatings.length > 0) {
              setOldRating(res.data.accomodationRatings[0]);
              setRating(res.data.accomodationRatings[0].rating);
            } else {
              setOldRating(null);
              setRating(0);
            }
          });
        return;
      }
      reservationService.changeReservationStatus(reservation).then((res) => {
        alert("Successfully canceled");
        setButtonClicked(!buttonClicked);
      });
    } else
      reservationService.deleteReservation(reservation.id).then((res) => {
        alert("Successfully deleted");
        setButtonClicked(!buttonClicked);
      });
  };

  const isInPast = (reservation: Reservation) => {
    const reservationEndDate = new Date(reservation.endDate.split(" ", 1)[0]);
    const currentDate = new Date();
    if (reservationEndDate.getTime() < currentDate.getTime()) {
      return true;
    }
    return false;
  };

  const handleRatingChange = (newRating: number) => {
    setRating(newRating);
  };

  const submitRating = () => {
    var date = formatDate(new Date());
    if (!oldRating) {
      var accomodationRating = new AccomodationRating({
        id: "",
        email: decodeToken()?.username,
        accomodationId: ratedAccomodation,
        rating: rating,
        date: date,
      });
      accomodationRatingService
        .createAccomodationRating(accomodationRating)
        .then((res) => {
          alert("Successfully rated!");
          setShowRating(false);
        });
    } else {
      var accomodationRating = new AccomodationRating({
        id: oldRating.id,
        email: decodeToken()?.username,
        accomodationId: ratedAccomodation,
        rating: rating,
        date: date,
      });
      accomodationRatingService
        .updateAccomodationRating(accomodationRating)
        .then((res) => {
          alert("Your rating for this accomodation is successfully changed!!");
          setShowRating(false);
        });
    }
  };

  function formatDate(date: Date) {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");

    return `${year}-${month}-${day}`;
  }

  const submitFlightsRecommendation = () => {
    var niz = selectedReservation.startDate.split(" ")
    var datum = niz[0]
    var niz2 = datum.split("-")
    var year = parseInt(niz2[0])
    var month = parseInt(niz2[1])-1
    var day = parseInt(niz2[2])

    var d = new Date(year, month, day,3,0,0)

    var searchCriteria1 = {
      startPlace: recommendParams.Departure,
      endPlace: "",
      numberOfPlaces: selectedReservation.guestNumber,
      date: d.toISOString() 
    }

    var nizz = selectedReservation.endDate.split(" ")
    var datumm = nizz[0]
    var nizz2 = datumm.split("-")
    var year2 = parseInt(nizz2[0])
    var month2 = parseInt(nizz2[1])-1
    var day2 = parseInt(nizz2[2])

    var d2 = new Date(year2, month2, day2,3,0,0)
    console.log("Arrival")
    console.log(recommendParams.Arrival)
    var searchCriteria2 = {
      startPlace: "", // gde je accomodation
      endPlace: recommendParams.Arrival,
      numberOfPlaces: selectedReservation.guestNumber,
      date: d2.toISOString() 
    }
    console.log("3")
    accomodationService.getAccomodation(selectedReservation.accomodationId).then((res) => {
      setAccomodationFromReservation(res.data.accomodation)
      searchCriteria1.endPlace = res.data.accomodation.location;
      searchCriteria2.startPlace = res.data.accomodation.location;
      flightService.getRecommendedFlights(searchCriteria1).then((res) => {
        setFlights1(res.data)
        console.log("flights1")
        console.log(flights1)
        flightService.getRecommendedFlights(searchCriteria2).then((res) => {
          setShowRecommendTables(true);
          setFlights2(res.data)
          console.log("flights2")
          console.log(flights2)
        });
      });
    });
  };

  const showRecomendFormFunc = (reservation: Reservation) => {
    
    setSelectedReservation(reservation);
    console.log(selectedReservation)
    setshowRecommendForm(true);
  };

  const buyTicket = (flight: Flight) => {
    /*var niz = flight.start.split("T")
    var datum = niz[0]
    var niz2 = datum.split("-")
    var year = parseInt(niz2[0])
    var month = parseInt(niz2[1])-1
    var day = parseInt(niz2[2])

    var satiNiz1 = niz[1].split(":")
    var hour = parseInt(satiNiz1[0])
    var min = parseInt(satiNiz1[1])
    
    var d = new Date(year, month, day,hour,min,0)

    var nizEnd = flight.end.split("T")
    var datumEnd = nizEnd[0]
    var nizEnd2 = datumEnd.split("-")
    var yearEnd = parseInt(nizEnd2[0])
    var monthEnd = parseInt(nizEnd2[1])-1
    var dayEnd = parseInt(nizEnd2[2])

    var satiNiz2 = nizEnd[1].split(":")
    var hourEnd = parseInt(satiNiz2[0])
    var minMin = parseInt(satiNiz2[1])
    
    var dEnd = new Date(yearEnd, monthEnd, dayEnd,hourEnd,minMin,0)*/
    console.log(flight)
    var dto = {apiKey: "",
                startTime: flight.start,
                startPlace: flight.startPlace,
                endTime: flight.end,
                endPlace: flight.endPlace,
                price: flight.pricePerPlace,
                count: selectedReservation.guestNumber,
                flightId: ""}
    userService.getUserByUsername().then((res) => {
      console.log(res.data)
      var user = res.data.user;
      console.log("user")
      console.log(user)
      dto.apiKey = res.data.user.apiKey;
      //dto.apiKey = "aaa"
      
      //console.log("api key")
      //console.log(dto.ApiKey)
      console.log("dto date 1")
      console.log(dto)
      ticketService.createTicket(dto).then((res) => {
        alert("Ticket bought!");
      });
    });
  };

  return (
    <>
      <div className="row">
        <div className="col-xl-8">
          <div className="card mb-4 mt-4">
            <div className="card-header">All reservations</div>
            <div className="card-body">
              <div className="table-responsive">
                <table className="table">
                  <thead>
                    <tr>
                      <th>Accomodation Id</th>
                      <th>Start date</th>
                      <th>End date</th>
                      <th>Status</th>
                      <th></th>
                    </tr>
                  </thead>
                  {reservations.map((reservation) => (
                    <tbody key={reservation.id}>
                      <tr>
                        <td>{reservation.accomodationId}</td>
                        <td>{reservation.startDate.split(" ", 1)}</td>
                        <td>{reservation.endDate.split(" ", 1)}</td>
                        <td>{reservation.status}</td>
                        <td>
                          {reservation.status === "REJECTED" ||
                          reservation.status === "CANCELED" ? (
                            ""
                          ) : (
                            <MDBBtn onClick={() => handleOnClick(reservation)}>
                              {reservation.status === "PENDING"
                                ? "DELETE"
                                : isInPast(reservation)
                                ? "RATE"
                                : "CANCEL"}
                            </MDBBtn>
                          )}
                        </td>
                        <td>
                          <MDBBtn onClick={() => showRecomendFormFunc(reservation)}>RECOMEND FLIGHTS</MDBBtn>
                        </td>
                      </tr>
                    </tbody>
                  ))}
                </table>
                {showRating && (
                  <div>
                    <h1>Rate the accomodation</h1>
                    {oldRating && <p>Old rating : {oldRating?.rating} </p>}
                    <StarRatings
                      rating={rating}
                      starRatedColor="#ffc107"
                      starEmptyColor="#ccc"
                      starHoverColor="#ffc107"
                      changeRating={handleRatingChange}
                      numberOfStars={5}
                      starDimension="30px"
                      starSpacing="5px"
                    />
                    <MDBBtn onClick={() => submitRating()}>SUBMIT</MDBBtn>
                  </div>
                )}

                {showRecommendForm && (
                  <div>
                    <MDBContainer fluid className="p-3 my-5 h-custom">
                      <MDBRow>
                      <MDBCol col="4" md="6">

                      <MDBInput
                        wrapperClass="mb-4"
                        label="Place before holiday"
                        onChange={(e) =>
                          setRecommendParams((prevState) => ({
                        ...prevState,
                        Departure: e.target.value,
                        }))
                        }
                        type="text"
                        size="lg"
                      />
                      </MDBCol>
                      <MDBCol>
                      <MDBInput
                        wrapperClass="mb-4"
                        label="Place after holiday"
                        onChange={(e) =>
                          setRecommendParams((prevState) => ({
                        ...prevState,
                        Arrival: e.target.value,
                        }))
                        }
                        type="text"
                        size="lg"
                      />

                      <div className="text-center text-md-start mt-4 pt-2">
                        <MDBBtn className="mb-0 px-5" size="lg" onClick={() => submitFlightsRecommendation()}>
                          Submit
                        </MDBBtn>
                      </div>
                    </MDBCol>
                  </MDBRow>
                  </MDBContainer>
                    
                  </div>
                )}
                {showRecommendTables && (
                  <div>
                    <table className="tableFlights">
                  <thead>
                    <tr>
                      <th>Start date</th>
                      <th>Start place</th>
                      <th>End date</th>
                      <th>End place</th>
                      <th>Remaining places</th>
                      <th>Price per place</th>
                      <th>Total price</th>
                      <th></th>
                    </tr>
                  </thead>
                  {flights1.map((flight) => (
                    <tbody key={flight.id}>
                      <tr>
                        <td>{flight.start}</td>
                        <td>{flight.startPlace}</td>
                        <td>{flight.end.split(" ", 1)}</td>
                        <td>{flight.endPlace}</td>
                        <td>{flight.remaining}</td>
                        <td>{flight.pricePerPlace}</td>
                        <td>{flight.pricePerPlace*selectedReservation.guestNumber}</td>
                        <td><MDBBtn onClick={() => buyTicket(flight)}>BUY</MDBBtn></td>
                      </tr>
                    </tbody>
                  ))}
                </table>
                <table className="tableFlights2">
                  <thead>
                    <tr>
                      <th>Start date</th>
                      <th>Start place</th>
                      <th>End date</th>
                      <th>End place</th>
                      <th>Remaining places</th>
                      <th>Price per place</th>
                      <th>Total price</th>
                      <th></th>
                    </tr>
                  </thead>
                  {flights2.map((flight) => (
                    <tbody key={flight.id}>
                      <tr>
                        <td>{flight.start}</td>
                        <td>{flight.startPlace}</td>
                        <td>{flight.end.split(" ", 1)}</td>
                        <td>{flight.endPlace}</td>
                        <td>{flight.remaining}</td>
                        <td>{flight.pricePerPlace}</td>
                        <td>{flight.pricePerPlace*selectedReservation.guestNumber}</td>
                        <td><MDBBtn onClick={() => buyTicket(flight)}>BUY</MDBBtn></td>
                      </tr>
                    </tbody>
                  ))}
                </table>
                  </div>
                )}
                
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
export default Reservations;

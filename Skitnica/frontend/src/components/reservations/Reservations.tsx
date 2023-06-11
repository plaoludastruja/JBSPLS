import { useEffect, useState } from "react";
import "./Reservations.css";
import reservationService from "../../services/reservation.service";
import Reservation from "../../model/Reservation";
import decodeToken from "../../services/auth.service";
import { MDBBtn } from "mdb-react-ui-kit";
import StarRatings from "react-star-ratings";
import accomodationRatingService from "../../services/accomodationRating.service";
import AccomodationRating from "../../model/AccomodationRating";

function Reservations() {
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [buttonClicked, setButtonClicked] = useState(false);
  const [rating, setRating] = useState(0);
  const [ratedAccomodation, setRatedAccomodation] = useState("");
  const [showRating, setShowRating] = useState(false);

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
      if (reservationEndDate.getTime() > currentDate.getTime()) {
        setShowRating(true);
        setRatedAccomodation(reservation.accomodationId);
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
    if (reservationEndDate.getTime() > currentDate.getTime()) {
      return true;
    }
    return false;
  };

  const handleRatingChange = (newRating: number) => {
    setRating(newRating);
  };

  const submitRating = () => {
    var date = formatDate(new Date());
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
  };

  function formatDate(date: Date) {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");

    return `${year}-${month}-${day}`;
  }
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
                      </tr>
                    </tbody>
                  ))}
                </table>
                {showRating && (
                  <div>
                    <h1>Rate the accomodation</h1>
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
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
export default Reservations;

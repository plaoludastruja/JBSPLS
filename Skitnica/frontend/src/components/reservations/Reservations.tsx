import { useEffect, useState } from "react";
import "./Reservations.css";
import reservationService from "../../services/reservation.service";
import Reservation from "../../interfaces/Reservation";
import decodeToken from "../../services/auth.service";

function Reservations() {
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [buttonClicked, setButtonClicked] = useState(false);

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
                      <th>Accomodation name</th>
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
                            <button onClick={() => handleOnClick(reservation)}>
                              {reservation.status === "PENDING"
                                ? "DELETE"
                                : "CANCEL"}
                            </button>
                          )}
                        </td>
                      </tr>
                    </tbody>
                  ))}
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
export default Reservations;

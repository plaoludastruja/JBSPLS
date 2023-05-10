import { useEffect, useState } from "react";
import "./Reservations.css";
import reservationService from "../../services/reservation.service";
import Reservation from "../../interfaces/Reservation";

function Reservations() {
  const [reservations, setReservations] = useState<Reservation[]>([]);

  useEffect(() => {
    reservationService.getReservations().then((response) => {
      setReservations(response.data.reservations);
    });
  }, [reservations]);

  const handleOnClick = (reservation: Reservation) => {
    if (reservation.status === "APPROVED")
      reservationService.changeReservationStatus(reservation).then((res) => {
        alert("Successfully cancelled");
      });
    else
      reservationService.deleteReservation(reservation.id).then((res) => {
        alert("Successfully deleted");
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
                      <th>Accomodation</th>
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
                        <td>{reservation.startDate}</td>
                        <td>{reservation.endDate}</td>
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

import { useEffect, useState } from "react";
import reservationService from "../../services/reservation.service";
import Reservation from "../../interfaces/Reservation";
import decodeToken from "../../services/auth.service";
import ReservationDto from "../../model/ReservationDto";

function HostsReservations() {
  const [reservations, setReservations] = useState<ReservationDto[]>([]);
  const [buttonClicked, setButtonClicked] = useState(false);
  const [all, setAll] = useState<Reservation[]>([]);

  useEffect(() => {
    reservationService
      .getPendingReservations(decodeToken()?.username)
      .then((response) => {
        setReservations(response.data.reservationDtos);
      });
  }, [buttonClicked]);

  const approve = (reservation: ReservationDto) => {
    
    reservation.startDate = reservation.startDate.split(" ", 1)[0];
    reservation.endDate = reservation.endDate.split(" ", 1)[0];
    reservationService.approve(reservation).then((res) => {
        alert("Successfully approved");
        setButtonClicked(!buttonClicked);
    });
  };

  const reject = (reservation: ReservationDto) => {
    reservation.startDate = reservation.startDate.split(" ", 1)[0];
    reservation.endDate = reservation.endDate.split(" ", 1)[0];
    reservationService.reject(reservation).then((res) => {
        alert("Successfully rejected");
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
                      <th>Guest username</th>
                      <th>Guests number</th>
                      <th>Number of cancelations</th>
                      <th></th>
                      <th></th>
                    </tr>
                  </thead>
                  {reservations.map((reservation) => (
                    <tbody key={reservation.id}>
                      <tr>
                        <td>{reservation.accomodationId}</td>
                        <td>{reservation.startDate.split(" ", 1)}</td>
                        <td>{reservation.endDate.split(" ", 1)}</td>
                        <td>{reservation.username}</td>
                        <td>{reservation.guestNumber}</td>
                        <td>{reservation.cancellationNum}</td>
                        <td><button onClick={() => approve(reservation)}>
                              APPROVE
                            </button></td>
                        <td><button onClick={() => reject(reservation)}>
                              REJECT
                            </button></td>
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
export default HostsReservations;
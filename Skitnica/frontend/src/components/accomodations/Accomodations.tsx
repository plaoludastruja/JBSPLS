import { useEffect, useState } from "react";
import Accomodation from "../../model/Accomodation";
import {
  MDBCard,
  MDBCardBody,
  MDBCardTitle,
  MDBCardText,
  MDBBtn,
} from "mdb-react-ui-kit";
import accomodationService from "../../services/accomodation.service";
import Appointment from "../../model/Appointment";
import priceService from "../../services/price.service";
import reservationService from "../../services/reservation.service";
import DateRange from "../../model/DateRange";
import decodeToken from "../../services/auth.service";
import "./Accomodations.css";

function Accomodations() {
  const [accomodations, setAccomodations] = useState<Accomodation[]>([]);
  const [appointments, setAppointments] = useState<Appointment[]>([]);
  const [appointment, setAppointment] = useState<Appointment>({
    id: "",
    accomodationId: "",
    start: "",
    end: "",
    priceType: "",
    price: 0,
    status: "Free",
  });

  const [appointmentsForCheck, setAppointmentsForCheck] = useState<
    Appointment[]
  >([]);
  const [isShown, setIsShown] = useState(false);
  const [appointmentForChange, setAppointmentForChange] = useState<Appointment>(
    {
      id: "",
      accomodationId: "",
      start: "",
      end: "",
      priceType: "",
      price: 0,
      status: "Free",
    }
  );

  useEffect(() => {
    accomodationService
      .getAccomodationsByHostUsername(decodeToken()?.username)
      .then((response) => {
        setAccomodations(response.data.accomodations);
      });
  }, []);

  const getByAccomodation = (accomodationId: string): void => {
    var appointments: Appointment[] = [];
    priceService.getByAccomodationId(accomodationId).then((response) => {
      setAppointments(response.data.appointments);
      console.log(accomodationId);
    });
    setIsShown(false);
  };

  const check = (appointmentChoosen: Appointment): void => {
    const dateRange: DateRange = {
      startDate: appointmentChoosen.start.split(" ", 1)[0],
      endDate: appointmentChoosen.end.split(" ", 1)[0],
      accomodationId: appointmentChoosen.accomodationId,
    };
    setAppointmentForChange(appointmentChoosen);
    console.log(dateRange);
    reservationService.check(dateRange).then((response) => {
      console.log(response.data.reservations.length);
      if (response.data.reservations.length == 0) {
        setIsShown(true);
      } else {
        setIsShown(false);
      }
    });
  };

  const editAppointment = (appointmentSent: Appointment): void => {
    console.log(appointmentSent);
    console.log(appointmentForChange);
    priceService.editAppointment(appointmentForChange).then((response) => {
      priceService
        .getByAccomodationId(appointmentForChange.accomodationId)
        .then((response) => {
          setAppointments(response.data.appointments);
        });
    });
  };

  return (
    <>
      Accomodations
      <div>
        {accomodations.map((accomodation, index) => (
          <div key={index}>
            <MDBCard>
              <MDBCardBody>
                <MDBCardTitle>{accomodation.name}</MDBCardTitle>
                <MDBCardText>
                  <div>
                    <img
                      src={`data:image/jpeg;base64,${accomodation.image}`}
                    ></img>
                    <div>{accomodation.location}</div>
                    <div>{accomodation.facilities}</div>
                    <div>
                      min: {accomodation.minNumberOfGuests}, max:{" "}
                      {accomodation.maxNumberOfGuests}
                    </div>
                  </div>
                </MDBCardText>
                <MDBBtn onClick={() => getByAccomodation(accomodation.id)}>
                  Show prices
                </MDBBtn>
              </MDBCardBody>
            </MDBCard>
          </div>
        ))}
        <div className="row">
          <div className="col-xl-8">
            <div className="card mb-4 mt-4">
              <div className="card-header">Prices</div>
              <div className="card-body">
                <div className="table-responsive">
                  <table className="table">
                    <thead>
                      <tr>
                        <th>Start date</th>
                        <th>End date</th>
                        <th>Type</th>
                        <th>Price</th>
                        <th></th>
                      </tr>
                    </thead>
                    {appointments.map((appointment) => (
                      <tbody key={appointment.id}>
                        <tr>
                          <td>{appointment.start.split(" ", 1)}</td>
                          <td>{appointment.end.split(" ", 1)}</td>
                          <td>{appointment.priceType}</td>
                          <td>{appointment.price}</td>
                          <td>
                            <MDBBtn onClick={() => check(appointment)}>
                              Edit
                            </MDBBtn>
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
      </div>
      <div>
        {isShown && (
          <div className="field">
            <label>
              Start:
              <input
                type="date"
                name="name"
                onChange={(e) =>
                  setAppointmentForChange((prevState) => ({
                    ...prevState,
                    start: e.target.value,
                  }))
                }
              />
            </label>
            <label>
              End:
              <input
                type="date"
                name="name"
                onChange={(e) =>
                  setAppointmentForChange((prevState) => ({
                    ...prevState,
                    end: e.target.value,
                  }))
                }
              />
            </label>
            <label>
              Price:
              <input
                type="number"
                name="name"
                onChange={(e) =>
                  setAppointmentForChange((prevState) => ({
                    ...prevState,
                    price: parseInt(e.target.value),
                  }))
                }
              />
            </label>
            <MDBBtn onClick={() => editAppointment(appointment)}>Edit</MDBBtn>
          </div>
        )}
      </div>
    </>
  );
}
export default Accomodations;

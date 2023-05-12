import { useEffect, useState } from "react";
import Accomodation from "../../model/Accomodation";
import { baseAxios } from "./../../services/api.service";
import {
    MDBCard,
    MDBCardBody,
    MDBCardTitle,
    MDBCardText,
    MDBBtn
  } from 'mdb-react-ui-kit';
import accomodationService from "../../services/accomodation.service";
import Appointment from "../../model/Appointment";
import priceService from "../../services/price.service";
import { response } from "express";

function Accomodations() {


  const [accomodations, setAccomodations] = useState<Accomodation[]>([]);
  const [appointments, setAppointments] = useState<Appointment[]>([]);

  useEffect(() => {
    accomodationService.getAccomodations().then((response) => {
      setAccomodations(response.data.accomodations);
    });
  }, [accomodations]);

  const getByAccomodation = (accomodationId: string) : void  => {
    var appointments : Appointment[] = [];
    priceService.getByAccomodationId(accomodationId).then((response) => {
      setAppointments(response.data.appointments);
      console.log(accomodationId)
    })
  }

  

  return (
    <>
      Accomodations
      <div>
      {accomodations.map((accomodation, index) => (
            <div>
                <MDBCard>
                <MDBCardBody>
                <MDBCardTitle>{accomodation.name}</MDBCardTitle>
                <MDBCardText>
                    <div>
                        <div key={index}>
                            {accomodation.location}
                        </div>
                        <div key={index + 1}>
                            {accomodation.facilities}
                        </div>
                        <div key={index + 2}>
                            min: {accomodation.minNumberOfGuests}, max: {accomodation.maxNumberOfGuests}
                        </div>
                    </div>
                </MDBCardText>
                <MDBBtn onClick={() => getByAccomodation(accomodation.id)}>Show prices</MDBBtn>
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
                    </tr>
                  </thead>
                  {appointments.map((appointment) => (
                    <tbody key={appointment.id}>
                      <tr>
                        <td>{appointment.start.split(" ", 1)}</td>
                        <td>{appointment.end.split(" ", 1)}</td>
                        <td>{appointment.priceType}</td>
                        <td>{appointment.price}</td>
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
    </>
  );
}
export default Accomodations;

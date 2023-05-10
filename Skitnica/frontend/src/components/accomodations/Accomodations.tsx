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

  useEffect(() => {
    accomodationService.getAccomodations().then((response) => {
      setAccomodations(response.data.accomodations);
    });
  }, [accomodations]);

  const getByAccomodation = (accomodationId: string) : Appointment[] => {
    var appointments : Appointment[] = [];
    priceService.getByAccomodationId(accomodationId).then((response) => {
        appointments = response.data.appointments;
    })
    return appointments;
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
                <MDBBtn>Button</MDBBtn>
                </MDBCardBody>
                </MDBCard>
            </div>
      ))}
      </div>
    </>
  );
}
export default Accomodations;

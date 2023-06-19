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
import accomodationRatingService from "../../services/accomodationRating.service";

function RecommendedAccomodations() {
  const [accomodations, setAccomodations] = useState<Accomodation[]>([]);
  const [accomodationsIds, setAccomodationsIds] = useState<string[]>([]);

  useEffect(() => {
    accomodationRatingService
      .getRecommendedAccomodations(decodeToken()?.username)
      .then((response) => {
        console.log("idovi")
        console.log(response.data.accommodations)
        response.data.accommodations.forEach((i) => {
          accomodationService.getAccomodation(i).then((response) => {
            var array = accomodations;
            array.push(response.data.accomodation);    
            setAccomodations(array)
          });
        })
        setAccomodationsIds(response.data.accommodations);
      });
  }, []);

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
              </MDBCardBody>
            </MDBCard>
          </div>
        ))}
       </div>
    </>
  );
}
export default RecommendedAccomodations;
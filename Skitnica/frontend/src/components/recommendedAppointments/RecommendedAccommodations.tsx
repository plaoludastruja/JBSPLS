import { useEffect, useState } from "react";
import Accomodation from "../../model/Accomodation";
import {
  MDBCard,
  MDBCardBody,
  MDBCardTitle,
  MDBCardText,
} from "mdb-react-ui-kit";
import accomodationService from "../../services/accomodation.service";
import accomodationRatingService from "../../services/accomodationRating.service";
import decodeToken from "../../services/auth.service";

function RecommendedAccomodations() {
  const [accomodations, setAccomodations] = useState<Accomodation[]>([]);
  const [accomodationsIds, setAccomodationsIds] = useState<string[]>([]);

  useEffect(() => {
    accomodationRatingService
      .getRecommendedAccomodations(decodeToken()?.username)
      .then((res) => {
        setAccomodationsIds(res.data.accommodations);
      });
  }, []);

  useEffect(() => {
    if (accomodationsIds.length > 0) {
      const getById = async (id: string) => {
        let res = await accomodationService.getAccomodation(id);
        setAccomodations((prevState) => [...prevState, res.data.accomodation]);
      };

      accomodationsIds.forEach(async (id: string) => await getById(id));
    }
  }, [accomodationsIds]);

  return (
    <>
      Accomodations
      <div>
        {accomodations.length > 0 &&
          accomodations.map((accomodation, index) => (
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

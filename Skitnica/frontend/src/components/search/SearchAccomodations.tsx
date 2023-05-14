import { useEffect, useState } from "react";
import {
  MDBCard,
  MDBCardBody,
  MDBCardTitle,
  MDBCardText,
  MDBBtn,
} from "mdb-react-ui-kit";
import accomodationService from "../../services/accomodation.service";
import reservationService from "../../services/reservation.service";
import SearchResult from "../../model/SearchResult";
import SearchParams from "../../model/SearchParams";
import Reservation from "../../interfaces/Reservation";
import decodeToken from "../../services/auth.service";

function SearchAccomodations() {
  const [searchResults, setSearchResults] = useState<SearchResult[]>([]);
  const [searchParams, setAppointmentForChange] = useState<SearchParams>({
    Location: "",
    GuestNumber: 0,
    StartDate: "",
    EndDate: "",
  });

  useEffect(() => {}, [searchResults]);

  const search = (): void => {
    accomodationService.searchAccomodations(searchParams).then((response) => {
      setSearchResults(response.data);
      console.log(response.data);
    });
  };

  const book = (searchResult: SearchResult): void => {
    console.log(searchResult.IsAutomaticApproved);
    var reservation = new Reservation({
      accomodationId: searchResult.AccomodationId,
      username: decodeToken()?.username,
      startDate: searchParams.StartDate,
      endDate: searchParams.EndDate,
      guestNumber: searchParams.GuestNumber,
      status:
        searchResult.IsAutomaticApproved === "true" ? "APPROVED" : "PENDING",
    });
    reservationService.createReservation(reservation).then((response) => {
      alert("Reservation is succesfully created");
    });
  };

  return (
    <>
      Accomodations
      <div className="field">
        <label>
          Start:
          <input
            type="date"
            name="name"
            onChange={(e) =>
              setAppointmentForChange((prevState) => ({
                ...prevState,
                StartDate: e.target.value,
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
                EndDate: e.target.value,
              }))
            }
          />
        </label>
        <label>
          Guest number:
          <input
            type="number"
            name="name"
            onChange={(e) =>
              setAppointmentForChange((prevState) => ({
                ...prevState,
                GuestNumber: parseInt(e.target.value),
              }))
            }
          />
        </label>
        <label>
          Location:
          <input
            type="text"
            name="name"
            onChange={(e) =>
              setAppointmentForChange((prevState) => ({
                ...prevState,
                Location: e.target.value,
              }))
            }
          />
        </label>
        <MDBBtn onClick={() => search()}>Search</MDBBtn>
      </div>
      <div>
        {searchResults.map((searchResult, index) => (
          <div key={index}>
            <MDBCard>
              <MDBCardBody>
                <MDBCardTitle>{searchResult.Name}</MDBCardTitle>
                <MDBCardText>
                  <div>
                    <div>{searchResult.Location}</div>
                    <div>{searchResult.Facilities}</div>
                    <div>
                      min: {searchResult.MinNumberOfGuests}, max:{" "}
                      {searchResult.MaxNumberOfGuests}
                    </div>
                    <div>{searchResult.TotalPrice}</div>
                    <button onClick={() => book(searchResult)}>Book now</button>
                  </div>
                </MDBCardText>
              </MDBCardBody>
            </MDBCard>
          </div>
        ))}
      </div>
      <div></div>
    </>
  );
}
export default SearchAccomodations;

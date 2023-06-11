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
import Reservation from "../../model/Reservation";
import decodeToken from "../../services/auth.service";
import accomodationRatingService from "../../services/accomodationRating.service";
import AccomodationRating from "../../model/AccomodationRating";

function SearchAccomodations() {
  const [searchResults, setSearchResults] = useState<SearchResult[]>([]);
  const [ratings, setRatings] = useState<AccomodationRating[]>([]);
  const [averageRating, setAverageRating] = useState(0);
  const [showRating, setShowRating] = useState(false);
  const [searchParams, setAppointmentForChange] = useState<SearchParams>({
    Location: "",
    GuestNumber: 0,
    StartDate: "",
    EndDate: "",
  });

  useEffect(() => {}, [searchResults]);

  const search = (): void => {
    if (
      searchParams.Location === "" ||
      searchParams.EndDate === "" ||
      searchParams.StartDate === "" ||
      searchParams.GuestNumber === 0
    ) {
      alert("Please enter all parameters");
    } else {
      accomodationService.searchAccomodations(searchParams).then((response) => {
        setSearchResults(response.data);
        console.log(response.data);
      });
    }
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
      hostUsername: searchResult.HostUsername,
    });
    reservationService.createReservation(reservation).then((response) => {
      alert("Reservation is succesfully created");
    });
  };

  const showRatings = (accomodation: SearchResult): void => {
    setShowRating(true);
    accomodationRatingService
      .getAccomodationRatingsByAccomodationId(accomodation.AccomodationId)
      .then((res) => {
        setRatings(res.data.accomodationRatings);
        var sum = 0;
        res.data.accomodationRatings.forEach((rating) => {
          sum += rating.rating;
        });
        setAverageRating(sum / res.data.accomodationRatings.length);
      });
  };

  const deleteRating = (rating: AccomodationRating): void => {
    accomodationRatingService
      .deleteAccomodationRating(rating.id)
      .then((res) => {
        alert("Your rating for this accomodation is deleted.");
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
                    <img
                      src={`data:image/jpeg;base64,${searchResult.Image}`}
                    ></img>
                    <div>Location: {searchResult.Location}</div>
                    <div>Facilities: {searchResult.Facilities}</div>
                    <div>
                      min: {searchResult.MinNumberOfGuests}, max:{" "}
                      {searchResult.MaxNumberOfGuests}
                    </div>
                    <div>Total price: {searchResult.TotalPrice}</div>
                    {searchResult.Prices.map((p, i) => (
                      <div key={i}>Unit price: {p}</div>
                    ))}
                    {searchResult.PriceType.map((pt, it) => (
                      <div key={it}>Price type: {pt}</div>
                    ))}
                    {decodeToken()?.role === "USER" && (
                      <MDBBtn onClick={() => book(searchResult)}>
                        Book now
                      </MDBBtn>
                    )}
                    <MDBBtn onClick={() => showRatings(searchResult)}>
                      Show ratings
                    </MDBBtn>
                  </div>
                </MDBCardText>
              </MDBCardBody>
            </MDBCard>
          </div>
        ))}
      </div>
      {showRating && (
        <div>
          <h1>Average Rating {averageRating}</h1>
          <div className="card-body">
            <div className="table-responsive">
              <table className="table">
                <thead>
                  <tr>
                    <th>Guest</th>
                    <th>Rating</th>
                    <th>Date</th>
                    <th></th>
                  </tr>
                </thead>
                {ratings.map((rating) => (
                  <tbody key={rating.id}>
                    <tr>
                      <td>{rating.email}</td>
                      <td>{rating.rating}</td>
                      <td>{rating.date}</td>
                      <td>
                        {rating.email === decodeToken()?.username && (
                          <MDBBtn onClick={() => deleteRating(rating)}>
                            DELETE
                          </MDBBtn>
                        )}
                      </td>
                    </tr>
                  </tbody>
                ))}
              </table>
            </div>
          </div>
        </div>
      )}
    </>
  );
}
export default SearchAccomodations;

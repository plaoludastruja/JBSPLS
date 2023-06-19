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
import FilterParams from "../../model/FilterParams";
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

  const [filterParams, setFilterParams] = useState<FilterParams>({
    Location: "",
    GuestNumber: 0,
    StartDate: "",
    EndDate: "",
    WiFi: "",
    Parking: "",
    AirCondtioning: "",
    BestHost: "false",
    MinPrice: "0",
    MaxPrice: "100000",
  });
  const [isChecked, setIsChecked] = useState(false);
  const [parking, setParking] = useState(false);
  const [airConditioning, setAirConditioning] = useState(false);

  const [bestHost, setBestHost] = useState(false);

  const [price1, setPrice1] = useState(false);
  const [price2, setPrice2] = useState(false);
  const [price3, setPrice3] = useState(false);

  const handleCheckboxChange = (event: { target: { checked: boolean | ((prevState: boolean) => boolean); }; }) => {
    setIsChecked(event.target.checked);
    //filter();
  };

  const handleCheckboxParkingChange = (event: { target: { checked: boolean | ((prevState: boolean) => boolean); }; }) => {
    setParking(event.target.checked);
    //filter();
  };

  const handleCheckboxAirConditioningChange = (event: { target: { checked: boolean | ((prevState: boolean) => boolean); }; }) => {
    setAirConditioning(event.target.checked);
    //filter();
  };

  const handleCheckboxBestHostChange = (event: { target: { checked: boolean | ((prevState: boolean) => boolean); }; }) => {
    setBestHost(event.target.checked);
    //filter();
  };

  const handleCheckboxPrice1Change = (event: { target: { checked: boolean | ((prevState: boolean) => boolean); }; }) => {
    setPrice1(event.target.checked);
    //filter();
  };

  const handleCheckboxPrice2Change = (event: { target: { checked: boolean | ((prevState: boolean) => boolean); }; }) => {
    setPrice2(event.target.checked);
    //filter();
  };

  const handleCheckboxPrice3Change = (event: { target: { checked: boolean | ((prevState: boolean) => boolean); }; }) => {
    setPrice3(event.target.checked);
    //filter();
  };

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
        filterParams.Location = searchParams.Location;
        filterParams.GuestNumber = searchParams.GuestNumber;
        filterParams.StartDate = searchParams.StartDate;
        filterParams.EndDate = searchParams.EndDate;
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

  const filter = (): void => {
    if (
      filterParams.Location === "" ||
      filterParams.EndDate === "" ||
      filterParams.StartDate === "" ||
      filterParams.GuestNumber === 0
    ) {
      alert("Please enter all parameters");
    } else {
      if(isChecked){
        filterParams.WiFi = "wifi"
      }else{
        filterParams.WiFi = ""
      }

      if(parking){
        filterParams.Parking = "parking"
      }else{
        filterParams.Parking = ""
      }

      if(airConditioning){
        filterParams.AirCondtioning = "air conditioning"
      }else{
        filterParams.AirCondtioning = ""
      }

      if(bestHost){
        filterParams.BestHost = "true"
      }else{
        filterParams.BestHost = "false"
      }

      if(price1){
        filterParams.MinPrice = "0"
        filterParams.MaxPrice = "100"
      }else if(price2){
        filterParams.MinPrice = "100"
        filterParams.MaxPrice = "200"
      }else if(price3){
        filterParams.MinPrice = "200"
        filterParams.MaxPrice = "100000"
      }else{
        filterParams.MinPrice = "0"
        filterParams.MaxPrice = "100000"
      }
      console.log("wifi")
      console.log(isChecked)
      console.log("parking")
      console.log(parking)
      console.log("air")
      console.log(airConditioning)

      console.log("price1")
      console.log(price1)
      console.log("price2")
      console.log(price2)
      console.log("price3")
      console.log(price3)

      console.log("host")
      console.log(bestHost)

      console.log("filter params")
      console.log(filterParams)

      accomodationService.filterAccomodations(filterParams).then((response) => {
        setSearchResults(response.data);
        console.log("filtrirao")
        console.log(response.data);
      });
    }
  }
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
        <h4>Filter</h4>
        
        <div>
          Facilities:
          <br></br>
          <label>
            <input
              type="checkbox"
              checked={isChecked}
              onChange={handleCheckboxChange}
            />
            WiFi
          </label>
          <br></br>
          <label>
            <input
              type="checkbox"
              checked={parking}
              onChange={handleCheckboxParkingChange}
            />
            Parking
          </label>
          <br></br>
          <label>
            <input
              type="checkbox"
              checked={airConditioning}
              onChange={handleCheckboxAirConditioningChange}
            />
            Air conditioning
          </label>
          <br></br>
          <br></br>
        </div>
        <div>
          Price:
          <br></br>
          <label>
            <input
              type="checkbox"
              checked={price1}
              onChange={handleCheckboxPrice1Change}
            />
            0-100
          </label>
          <br></br>
          <label>
            <input
              type="checkbox"
              checked={price2}
              onChange={handleCheckboxPrice2Change}
            />
            100-200
          </label>
          <br></br>
          <label>
            <input
              type="checkbox"
              checked={price3}
              onChange={handleCheckboxPrice3Change}
            />
            200+
          </label>
          <br></br>
          <br></br>
        </div>
        <div>
          Host mark
          <br></br>
          <label>
            <input
              type="checkbox"
              checked={bestHost}
              onChange={handleCheckboxBestHostChange}
            />
            Best marks
          </label>
          <br></br>
          <MDBBtn onClick={() => filter()}>Filter</MDBBtn>
        </div>
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

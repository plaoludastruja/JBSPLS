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
import reservationService from "../../services/reservation.service";
import DateRange from "../../model/DateRange";
import SearchResult from "../../model/SearchResult";
import SearchParams from "../../model/SearchParams";

function SearchAccomodations() {


  const [searchResults, setSearchResults] = useState<SearchResult[]>([]);
  const [searchParams, setAppointmentForChange] = useState<SearchParams>({
    Location: "",
    GuestNumber: 0,
    StartDate: "",
    EndDate: "",
  });


  useEffect(() => {
    /*accomodationService.searchAccomodations().then((response) => {
        setSearchResults(response.data.accomodations);
    });*/
  }, [searchResults]);

  const search = () : void  => {
    accomodationService.searchAccomodations(searchParams).then((response) => {
      setSearchResults(response.data);
      console.log(response.data);
    })
     
  }

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
                        <div>
                            {searchResult.Location}
                        </div>
                        <div>
                            {searchResult.Facilities}
                        </div>
                        <div>
                            min: {searchResult.MinNumberOfGuests}, max: {searchResult.MaxNumberOfGuests}
                        </div>
                        <div>
                            {searchResult.TotalPrice}
                        </div>
                    </div>
                </MDBCardText>
                </MDBCardBody>
                </MDBCard>
            </div>
            
      ))}
      </div>
      <div>
      
      </div>
    </>
  );
}
export default SearchAccomodations;
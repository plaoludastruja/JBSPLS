import Accomodation from "../model/Accomodation";
import AccomodationResponse from "../model/AccomodationResponse";
import AccomodationsResponse from "../model/AccomodationsResponse";
import FilterParams from "../model/FilterParams";
import SearchParams from "../model/SearchParams";
import SearchResult from "../model/SearchResult";
import { baseAxios } from "./api.service";

const createAccomodation = (accomodation: Accomodation)  => baseAxios.post('/accomodation', accomodation)
const getAccomodations = async()  => await baseAxios.get<AccomodationsResponse>('/accomodation')
const getAccomodationsByHostUsername = (hostUsername: string|undefined)  =>  baseAxios.get<AccomodationsResponse>('/accomodation/host/' + hostUsername)
const searchAccomodations = async(searchParams: SearchParams)  => await baseAxios.get<SearchResult[]>(`/accomodation/search/${searchParams.Location}/${searchParams.GuestNumber}/${searchParams.StartDate.split("-")[2]}/${searchParams.StartDate.split("-")[1]}/${searchParams.StartDate.split("-")[0]}/${searchParams.EndDate.split("-")[2]}/${searchParams.EndDate.split("-")[1]}/${searchParams.EndDate.split("-")[0]}`)
const getAccomodationById = (id: number) => baseAxios.get<AccomodationsResponse>('/accomodation/' + id)
const filterAccomodations = async(filterParams: FilterParams)  => await baseAxios.get<SearchResult[]>(`/accomodation/filter/${filterParams.Location}/${filterParams.GuestNumber}/${filterParams.StartDate.split("-")[2]}/${filterParams.StartDate.split("-")[1]}/${filterParams.StartDate.split("-")[0]}/${filterParams.EndDate.split("-")[2]}/${filterParams.EndDate.split("-")[1]}/${filterParams.EndDate.split("-")[0]}/${filterParams.WiFi}/${filterParams.Parking}/${filterParams.AirCondtioning}/${filterParams.BestHost}/${filterParams.MinPrice}/${filterParams.MaxPrice}`)
const getAccomodation = (id: string) => baseAxios.get<AccomodationResponse>('/accomodation/' + id)

export default {
    createAccomodation,
    getAccomodations,
    getAccomodationsByHostUsername,
    searchAccomodations,
    filterAccomodations,
    getAccomodation
  }
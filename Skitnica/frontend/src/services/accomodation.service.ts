import Accomodation from "../model/Accomodation";
import AccomodationsResponse from "../model/AccomodationsResponse";
import SearchParams from "../model/SearchParams";
import SearchResult from "../model/SearchResult";
import { baseAxios } from "./api.service";

const createAccomodation = (accomodation: Accomodation)  => baseAxios.post('/accomodation', accomodation)
const getAccomodations = async()  => await baseAxios.get<AccomodationsResponse>('/accomodation')
const getAccomodationsByHostUsername = (hostUsername: string|undefined)  =>  baseAxios.get<AccomodationsResponse>('/accomodation/host/' + hostUsername)
const searchAccomodations = async(searchParams: SearchParams)  => await baseAxios.get<SearchResult[]>(`/accomodation/search/${searchParams.Location}/${searchParams.GuestNumber}/${searchParams.StartDate.split("-")[2]}/${searchParams.StartDate.split("-")[1]}/${searchParams.StartDate.split("-")[0]}/${searchParams.EndDate.split("-")[2]}/${searchParams.EndDate.split("-")[1]}/${searchParams.EndDate.split("-")[0]}`)
const getAccomodationById = (id: number) => baseAxios.get<AccomodationsResponse>('/accomodation/' + id)


export default {
    createAccomodation,
    getAccomodations,
    getAccomodationsByHostUsername,
    searchAccomodations
  }
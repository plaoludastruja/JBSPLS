import Accomodation from "../model/Accomodation";
import AccomodationsResponse from "../model/AccomodationsResponse";
import { baseAxios } from "./api.service";

const createAccomodation = (accomodation: Accomodation)  => baseAxios.post('/accomodation', accomodation)
const getAccomodations = async()  => await baseAxios.get<AccomodationsResponse>('/accomodation')
//const getAccomodationById = async(id : string)  => await baseAxios.get<Accomodation>('/accomodation' + id)


export default {
    createAccomodation,
    getAccomodations,
  }
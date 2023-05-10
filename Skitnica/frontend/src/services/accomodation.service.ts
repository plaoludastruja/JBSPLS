import Accomodation from "../model/Accomodation";
import AccomodationResponse from "../model/AccomodationResponse";
import { baseAxios } from "./api.service";

const createAccomodation = (accomodation: Accomodation)  => baseAxios.post('/accomodation', accomodation)
const getAccomodations = async()  => await baseAxios.get<AccomodationResponse>('/accomodation')

export default {
    createAccomodation,
    getAccomodations
  }
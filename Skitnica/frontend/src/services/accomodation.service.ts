import Accomodation from "../model/Accomodation";
import AccomodationsResponse from "../model/AccomodationsResponse";
import { baseAxios } from "./api.service";

const createAccomodation = (accomodation: Accomodation)  => baseAxios.post('/accomodation', accomodation)
const getAccomodations = async()  => await baseAxios.get<AccomodationsResponse>('/accomodation')


export default {
    createAccomodation,
    getAccomodations,
  }
import Accomodation from "../model/Accomodation";
import { baseAxios } from "./api.service";

const createAccomodation = (accomodation: Accomodation)  => baseAxios.post('/accomodation', accomodation)
const getAccomodations = async()  => await baseAxios.get<Accomodation[]>('/accomodation')

export default {
    createAccomodation,
    getAccomodations
  }
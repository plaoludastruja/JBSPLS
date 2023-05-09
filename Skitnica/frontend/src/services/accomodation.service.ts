import Accomodation from "../model/Accomodation";
import { baseAxios } from "./api.service";

const createAccomodation = (accomodation: Accomodation)  => baseAxios.post('/accomodation', accomodation)


export default {
    createAccomodation
  }
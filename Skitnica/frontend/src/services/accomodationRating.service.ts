import AccomodationRating from "../model/AccomodationRating"
import AccomodationRatingsResponse from "../model/AccomodationRatingsResponse"
import { baseAxios } from "./api.service"

const createAccomodationRating = (accomodationRating: AccomodationRating)  => baseAxios.post('/accomodation-rating', accomodationRating)
const getAccomodationRatingsByAccomodationId = async(accomodationId: string)  => await baseAxios.get<AccomodationRatingsResponse>('/accomodation-rating/' + accomodationId)
const deleteAccomodationRating = (accomodationRatingId : string)  => baseAxios.delete('/accomodation-rating/' + accomodationRatingId)

export default {
    createAccomodationRating,
    getAccomodationRatingsByAccomodationId,
    deleteAccomodationRating
  }
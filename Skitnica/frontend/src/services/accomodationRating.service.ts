import AccomodationRating from "../model/AccomodationRating"
import AccomodationRatingsResponse from "../model/AccomodationRatingsResponse"
import GetAllRecommendedResponse from "../model/GetAllRecommendedResponse"
import { baseAxios } from "./api.service"

const createAccomodationRating = (accomodationRating: AccomodationRating)  => baseAxios.post('/accomodation-rating', accomodationRating)
const getAccomodationRatingsByAccomodationId = async(accomodationId: string)  => await baseAxios.get<AccomodationRatingsResponse>('/accomodation-rating/' + accomodationId)
const deleteAccomodationRating = (accomodationRatingId : string)  => baseAxios.delete('/accomodation-rating/' + accomodationRatingId)
const getAccomodationRatingsByAccomodationAndUser = async(accomodationId: string, email: string | undefined)  => await baseAxios.get<AccomodationRatingsResponse>('/accomodation-rating/' + accomodationId + '/' + email)
const updateAccomodationRating = (accomodationRating: AccomodationRating)  => baseAxios.put('/accomodation-rating', accomodationRating)
const getRecommendedAccomodations = async(email: string| undefined)  => await baseAxios.get<GetAllRecommendedResponse>('/accomodation-rating/recommended/' + email)

export default {
    createAccomodationRating,
    getAccomodationRatingsByAccomodationId,
    deleteAccomodationRating,
    getAccomodationRatingsByAccomodationAndUser,
    updateAccomodationRating,
    getRecommendedAccomodations
  }
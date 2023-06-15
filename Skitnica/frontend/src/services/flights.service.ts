import { Flight } from "../model/Flight";
import RecomendFlightsParam from "../model/RecomendFlightsParam";
import SearchFlightsDto from "../model/SearchFlightsDto";
import { baseAxios, baseAxios2 } from "./api.service";

const getRecommendedFlights = async (dto: SearchFlightsDto) => await baseAxios2.post<Flight[]>('/flight/search', dto)

export default {
    getRecommendedFlights
  }
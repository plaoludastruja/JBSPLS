import { Flight } from "../model/Flight";
import RecomendFlightsParam from "../model/RecomendFlightsParam";
import SearchFlightsDto from "../model/SearchFlightsDto";
import TicketSkitnicaDTO from "../model/TicketSkitnicaDTO";
import { baseAxios, baseAxios2 } from "./api.service";

const createTicket = async (dto: any) => await baseAxios2.post<any>('/ticket/skitnica', dto)

export default {
    createTicket
  }
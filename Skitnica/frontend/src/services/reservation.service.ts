import Reservation from "../interfaces/Reservation";
import ReservationsResponse from "../interfaces/ReservationsResponse";
import { baseAxios } from "./api.service";

const getReservations = (email: any)  => baseAxios.get<ReservationsResponse>('/reservation/all/' + email)
const changeReservationStatus = (reservation : Reservation)  => baseAxios.put('/reservation', reservation)
const deleteReservation = (reservationId : string)  => baseAxios.delete('/reservation/' + reservationId)

export default {
    getReservations,
    changeReservationStatus,
    deleteReservation
  }
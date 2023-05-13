import Reservation from "../interfaces/Reservation";
import ReservationsResponse from "../interfaces/ReservationsResponse";
import DateRange from "../model/DateRange";
import { baseAxios } from "./api.service";

const getReservations = (email: any)  => baseAxios.get<ReservationsResponse>('/reservation/all/' + email)
const changeReservationStatus = (reservation : Reservation)  => baseAxios.put('/reservation', reservation)
const deleteReservation = (reservationId : string)  => baseAxios.delete('/reservation/' + reservationId)
const check = (dateRange: DateRange)  => baseAxios.post<ReservationsResponse>('/reservation/check', dateRange)
const createReservation = (reservation: Reservation)  => baseAxios.post('/reservation', reservation)

export default {
    getReservations,
    changeReservationStatus,
    deleteReservation,
    check,
    createReservation
  }
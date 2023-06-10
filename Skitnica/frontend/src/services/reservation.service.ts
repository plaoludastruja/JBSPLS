import Reservation from "../model/Reservation";
import ReservationsResponse from "../model/ReservationsResponse";
import DateRange from "../model/DateRange";
import ReservationDto from "../model/ReservationDto";
import ReservationDtosResponse from "../model/ReservationDtosResponse";
import { baseAxios } from "./api.service";
import HostsForGuestResponse from "../model/HostsForGuest";

const getReservations = (email: any)  => baseAxios.get<ReservationsResponse>('/reservation/all/' + email)
const changeReservationStatus = (reservation : Reservation)  => baseAxios.put('/reservation', reservation)
const deleteReservation = (reservationId : string)  => baseAxios.delete('/reservation/' + reservationId)
const check = (dateRange: DateRange)  => baseAxios.post<ReservationsResponse>('/reservation/check', dateRange)
const getPendingReservations = (hostUsername: any)  => baseAxios.get<ReservationDtosResponse>('/reservation/allPending/' + hostUsername)
const approve = (reservationDto : ReservationDto)  => baseAxios.put('/reservation/approve', reservationDto)
const reject = (reservationDto : ReservationDto)  => baseAxios.put('/reservation/reject', reservationDto)
const createReservation = (reservation: Reservation)  => baseAxios.post('/reservation', reservation)
const getByGuest = (guestUsername: string|undefined) => baseAxios.get<HostsForGuestResponse>('/reservation/by-guest/' + guestUsername)

export default {
    getReservations,
    changeReservationStatus,
    deleteReservation,
    check,
    getPendingReservations,
    approve,
    reject,
    createReservation,
    getByGuest
  }
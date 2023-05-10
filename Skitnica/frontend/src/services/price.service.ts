import Appointment from "../model/Appointment"
import AppointmentsResponse from "../model/AppointmentsResponse"
import { baseAxios } from "./api.service"

const createAppointment = (appointment: Appointment)  => baseAxios.post('/appointment', appointment)
const getByAccomodationId =  async(accomodationId: string)  => await baseAxios.get<AppointmentsResponse>('/appointment/${accomodationId}')

export default {
    createAppointment,
    getByAccomodationId
  }
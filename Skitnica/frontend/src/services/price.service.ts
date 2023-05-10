import Appointment from "../model/Appointment"
import { baseAxios } from "./api.service"

const createAppointment = (appointment: Appointment)  => baseAxios.post('/appointment', appointment)

export default {
    createAppointment
  }
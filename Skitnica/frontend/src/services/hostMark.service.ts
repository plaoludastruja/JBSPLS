import HostMarksResponse from "../model/HostMarkResponse";
import { baseAxios } from "./api.service";

const getHostMarks = async()  => await baseAxios.get<HostMarksResponse>('/hostMark')

export default {
    getHostMarks
  }
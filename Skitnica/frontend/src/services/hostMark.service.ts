import HostMark from "../model/HostMark";
import HostMarksResponse from "../model/HostMarkResponse";
import { baseAxios } from "./api.service";

const getHostMarks = async()  => await baseAxios.get<HostMarksResponse>('/hostMark')
const getByHostAndUsername =  async(username: string|undefined, hostUsername: string)  => await baseAxios.get<HostMarksResponse>(`/hostmark/${username}/${hostUsername}`)
const createHostGrade = (hostMark: HostMark)  => baseAxios.post('/hostmark', hostMark)

export default {
    getHostMarks,
    getByHostAndUsername,
    createHostGrade
  }
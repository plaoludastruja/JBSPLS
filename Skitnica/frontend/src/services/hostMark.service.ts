import HostMark from "../model/HostMark";
import HostMarksResponse from "../model/HostMarkResponse";
import { baseAxios } from "./api.service";

const getHostMarks = async()  => await baseAxios.get<HostMarksResponse>('/hostMark')
const getByHostAndUsername =  async(username: string|undefined, hostUsername: string)  => await baseAxios.get<HostMarksResponse>(`/hostmark/${username}/${hostUsername}`)
const createHostGrade = (hostMark: HostMark)  => baseAxios.post('/hostmark', hostMark)
const editHostGrade = (hostMark: HostMark)  => baseAxios.put('/hostmark', hostMark)
const deleteHostGrade = (id: string)  => baseAxios.delete(`/hostmark/${id}`)
const getByHost =  async(username: string)  => await baseAxios.get<HostMarksResponse>(`/hostmark/host/${username}`)

export default {
    getHostMarks,
    getByHostAndUsername,
    createHostGrade,
    editHostGrade,
    deleteHostGrade,
    getByHost
  }
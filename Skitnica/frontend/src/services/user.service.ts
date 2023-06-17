import LoginDTO from "../model/LoginDTO";
import User from "../model/User";
import UsersResponse from "../model/UsersResponse";
import USersResponse from "../model/UsersResponse";
import { baseAxios } from "./api.service";
import { setToken } from "./token.service";

const getUserByUsername = () => baseAxios.get('/user/getByUsername')
const editUser = (user: User) => baseAxios.put('/user', user)
const registerUser = (user: User) => baseAxios.post('/user', user)
const loginUser = (loginDTO: LoginDTO) => baseAxios.post('user/login', loginDTO)
const deleteUser = (userId: string) => baseAxios.delete('/user/' + userId)

const getAllUnread = (username: string) => baseAxios.get('/notificatin/getAllUnread/' + username)
const getAllByReceiver = (username: string) => baseAxios.get('/notification/getByReceiver/' + username)
const readAllByUsername = (username: string) => baseAxios.put('/notification/readAllByUsername/' + username)

const isBestHost = (hostUsername: string) => baseAxios.get('/isBestHost/' + hostUsername)
const getHosts = () => baseAxios.get<UsersResponse>('/user/hosts')

export default {
    registerUser,
    loginUser,
    getUserByUsername,
    editUser,
    deleteUser,
    isBestHost,
    getHosts,
    getAllUnread,
    getAllByReceiver,
    readAllByUsername
}
import { useNavigate } from "react-router-dom";
import LoginDTO from "../model/LoginDTO";
import User from "../model/User";
import { baseAxios } from "./api.service";
import { setToken } from "./token.service";

const getUserByUsername = () => baseAxios.get('/user/getByUsername')
const editUser = (user: User) => baseAxios.put('/user', user)
const registerUser = (user: User) => baseAxios.post('/user', user)
const loginUser = (loginDTO: LoginDTO) => baseAxios.post('user/login', loginDTO)
const deleteUser = (userId: string) => baseAxios.delete('/user/' + userId)

export default {
    registerUser,
    loginUser,
    getUserByUsername,
    editUser,
    deleteUser
}
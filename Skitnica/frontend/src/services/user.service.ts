import { useNavigate } from "react-router-dom";
import LoginDTO from "../model/LoginDTO";
import User from "../model/User";
import { baseAxios } from "./api.service";
import { setToken } from "./token.service";


const registerUser = (user: User) => baseAxios.post('/user', user)
const loginUser = (loginDTO: LoginDTO) => baseAxios.post('user/login', loginDTO).then((res) => {
    setToken(res.data.token)
})

export default {
    registerUser,
    loginUser
}
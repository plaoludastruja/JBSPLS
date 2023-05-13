import jwtDecode from "jwt-decode"
import { Jwt } from "../model/Jwt"

const decodeToken = (): Jwt | undefined => {
    const token = localStorage.getItem('token')
    if(token) 
      return jwtDecode<Jwt>(token)
    return undefined
  }

export default decodeToken;

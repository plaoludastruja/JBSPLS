import axios from 'axios'

import { getToken } from './token.service'

export const baseAxios = axios.create({ baseURL: "http://localhost:8000/" })
baseAxios.interceptors.request.use(
  config => {
    const token = getToken()
    if (token) {
      config.headers['Authorization'] = 'Bearer ' + token
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

export const baseAxios2 = axios.create({ baseURL: "http://localhost:8080/" })
baseAxios2.interceptors.request.use(
  config => {
    const token = getToken()
    if (token) {
      config.headers['Authorization'] = 'Bearer ' + token
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

import axios from "axios"
import {ElMessage} from "element-plus"

export const api = axios.create({
  baseURL: 'https://axogc.net:8080',
})

export function request<T = any>(method: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE', url: string, data?: any): Promise<T> {
  return api.request({
    method, url,
    params: ['GET', 'DELETE'].includes(method) ? data : undefined,
    data: ['POST', 'PUT', 'PATCH'].includes(method) ? data : undefined,
  })
}

api.interceptors.response.use(res => {
  if (res.data.message) {
    ElMessage({type: 'info', message: res.data.message})
  }
  return res.data.data
}, async err => {
  const res = err.response
  ElMessage({type: 'error', message: res.data.message})
  return Promise.reject(err)
})


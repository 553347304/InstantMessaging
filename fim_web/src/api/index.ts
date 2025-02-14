import axios from "axios";
import {ElMessage} from "element-plus";
import {useStore} from "@/stores/stores.ts";

export const useAxios = axios.create({
    baseURL: "",
    timeout: 3000,
})

// 请求拦截器
useAxios.interceptors.request.use((config) => {
    const store = useStore()
    config.headers["token"] = store.userInfo.token
    return config
})

// 响应拦截器
useAxios.interceptors.response.use((response) => {
    if (response.status !== 200) {
        console.log("请求错误", response.status)
        ElMessage.error(response.statusText)
        return Promise.reject(response.statusText)
    }
    return response.data
}, (err) => {
    console.log("服务错误", err)
    ElMessage.error(err.message)
    return Promise.reject(err.message)
})
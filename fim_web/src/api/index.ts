import axios from "axios";
import {ElMessage} from "element-plus";
import {useStore} from "@/stores/stores.ts";
// import {useStore} from "@/stores";

export const useAxios = axios.create({
    baseURL: "", //
    timeout: 3000,
})
//
// // 请求拦截器
// useAxios.interceptors.request.use((config) => {
//     const store = useStore()
//     config.headers["token"] = store.userInfo.token
//     return config
// })


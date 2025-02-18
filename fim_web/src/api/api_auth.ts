import {type baseResponse, baseURL} from "@/api/response.ts";
import {useAxios} from "@/api/index.ts";


export namespace typeAuth {
    export interface userInfo {
        username: string
        role: number
        user_id: number
        token: string
    }

    export const userInfo = (): userInfo => ({
        username: "",
        role: 0,
        user_id: 0,
        token: "",
    })
    
    export interface loginRequest {
        username: string
        password: string
    }

    export interface loginResponse {
        token: string
    }

    export const loginRequest = (): loginRequest => ({
        username: "",
        password: "",
    })

    export interface openLoginResponse {
        code: string
        flag: string
    }
    export const openLoginResponse = (): openLoginResponse => ({
        code: "",
        flag: "",
    })

}

class m_api {
    Login = (data: typeAuth.loginRequest): Promise<baseResponse<typeAuth.loginResponse>> => {
        return useAxios.post(baseURL.auth.login, data)
    }
    OpenLogin = (data: typeAuth.openLoginResponse): Promise<baseResponse<typeAuth.loginResponse>> => {
        return useAxios.post(baseURL.auth.open_login, data)
    }
}

export const ApiAuth = new m_api()
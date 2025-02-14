import {type baseResponse, baseURL} from "@/api/response.ts";
import {useAxios} from "@/api/index.ts";


export namespace typeAuth {
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
}

class api {
    Login = (data: typeAuth.loginRequest): Promise<baseResponse<typeAuth.loginResponse>> => {
        return useAxios.post(baseURL.auth.login, data)
    }
}

export const ApiAuth = new api()
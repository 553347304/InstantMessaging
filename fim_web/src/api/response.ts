import {ElMessage} from "element-plus";


export interface baseResponse<T> {
    code: number
    message: string
    data: T
}

export interface listResponse<T> {
    total: number
    list: T[]
}

export function baseResponse(response: baseResponse<any>): boolean {
    if (response.code) {
        ElMessage.error(response.message);
        return false;
    }
    return true;
}

export const baseURL = {
    auth: {
        login: "/api/auth/login",
        open_login: "/api/auth/open_login",
    },
    user: {
        info: "/api/user/user_info",
        update: "/api/user/user_info",
    },
    setting: {
        info: "/api/setting/info",
    },
    file: {
        upload: "/api/file/upload",
    },
}
export interface baseResponse<T> {
    code: number
    message: string
    data: T
}

export interface listResponse<T> {
    total: number
    list: T[]
}

export const baseURL = {
    auth: {
        login: "/api/auth/login",
    },
    user: {
        info: "/api/user/user_info",
    },
}
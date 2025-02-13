

export interface baseResponse<T> {
    code: number
    message: string
    data: T
}

export interface listResponse<T> {
    total: number
    list: T[]
}
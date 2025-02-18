import type {baseResponse} from "@/api/response.ts";
import {useAxios} from "@/api/index.ts";
import {baseURL} from "@/api/response.ts";

export namespace typeUser {
    export interface infoResponse {
        id: number
        created_at: string
        updated_at: string
        username: string
        password: string
        sign: string
        avatar: string
        ip: string
        addr: string
        role: number
        open_id: string
        register_source: string
        user_config_model: {
            id: number
            created_at: string
            updated_at: string
            user_id: number
            recall_message: string
            friend_online: boolean
            sound: boolean
            secure_link: boolean
            save_password: boolean
            search_user: number
            valid: number
            valid_info: {
                issue: string
                answer: string
            }
            online: boolean
            curtail_chat: boolean
            curtail_add_user: boolean
            curtail_create_group: boolean
            curtail_add_group: boolean
        }
        top: {
            group_id: []
        }
    }

    export const infoResponse = (): infoResponse => ({
        id: 0,
        created_at: "",
        updated_at: "",
        username: "",
        password: "",
        sign: "",
        avatar: "",
        ip: "",
        addr: "",
        role: 0,
        open_id: "",
        register_source: "",
        user_config_model: {
            id: 0,
            created_at: "",
            updated_at: "",
            user_id: 0,
            recall_message: "",
            friend_online: true,
            sound: false,
            secure_link: false,
            save_password: false,
            search_user: 0,
            valid: 0,
            valid_info: {
                issue: "",
                answer: "",
            },
            online: false,
            curtail_chat: false,
            curtail_add_user: false,
            curtail_create_group: true,
            curtail_add_group: false,
        },
        top: {
            group_id: [],
        },
    })

    export interface updateRequest {
        user_info?: {
            username?: string
            sign?: string
            avatar?: string
        }
        user_config?: {
            recall_message?: string
            friend_online?: boolean
            sound?: boolean
            secure_link?: boolean
            save_password?: boolean
            search_user?: number
            valid?: number
            valid_info?: {
                issue?: string
                answer?: string
            }
        }
    }

    export const updateRequest = (): updateRequest => ({
        user_info: {
            username: "",
            sign: "",
            avatar: "",
        },
        user_config: {
            recall_message: "",
            friend_online: true,
            sound: true,
            secure_link: true,
            save_password: true,
            search_user: 0,
            valid: 0,
            valid_info: {
                issue: "",
                answer: "",
            },
        },
    })


}

class m_api {
    Info = (): Promise<baseResponse<typeUser.infoResponse>> => {
        return useAxios.get(baseURL.user.info)
    }
    Update = (data: typeUser.updateRequest): Promise<baseResponse<string>> => {
        return useAxios.put(baseURL.user.update, data)
    }
}

export const ApiUser = new m_api()
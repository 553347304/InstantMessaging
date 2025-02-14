export namespace typeUser {
    export interface userInfo {
        username: string
        role: number
        user_id: number
        token: string
    }

    export interface userInfoResponse {
        user_id: number
        name: string
        sign: string
        avatar: string
        recall_message: string
        friend_online: boolean
        sound: boolean
        secure_link: boolean
        save_password: boolean
        search_user: number
        valid: number
        valid_info: {
            issue: []
            answer: []
        }
    }


    export const userInfo = (): userInfo => ({
        username: "",
        role: 0,
        user_id: 0,
        token: "",
    })
}

class api {

}

export const ApiUser = new api()
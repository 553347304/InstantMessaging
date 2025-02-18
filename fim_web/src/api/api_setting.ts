import type {baseResponse} from "@/api/response.ts";
import {useAxios} from "@/api/index.ts";
import {baseURL} from "@/api/response.ts";

export namespace typeSetting {
    export interface infoResponse {
        id: number
        site: {
            created_at: string
            bei_an: string
            version: string
            image_qq: string
            image_wechat: string
            url_bili_bili: string
            url_gitee: string
            url_github: string
        }
        open_login: {
            qq: {
                enable: boolean
                app_id: string
                key: string
                redirect: string
                webPath: string
            }
        }
    }

    export const infoResponse = (): infoResponse => ({
        id: 0,
        site: {
            created_at: "",
            bei_an: "",
            version: "",
            image_qq: "",
            image_wechat: "",
            url_bili_bili: "",
            url_gitee: "",
            url_github: "",
        },
        open_login: {
            qq: {
                enable: true,
                app_id: "",
                key: "",
                redirect: "",
                webPath: "",
            },
        },
    })


}


class m_api {
    Info = (): Promise<baseResponse<typeSetting.infoResponse>> => {
        return useAxios.get(baseURL.setting.info)
    }
}

export const ApiSetting = new m_api()
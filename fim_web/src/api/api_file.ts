import type {baseResponse} from "@/api/response.ts";
import {useAxios} from "@/api/index.ts";
import {baseURL} from "@/api/response.ts";

export namespace typeFile {
    export interface uploadResponse {
        url: string
    }


}


class m_api {
    Upload = (file: File, type: "avatar" | "group_avatar" | "chat"): Promise<baseResponse<typeFile.uploadResponse>> => {
        const form = new FormData()
        form.set("file", file)
        form.set("type", type)
        return useAxios.post(baseURL.file.upload, form, {
            headers: {
                "Content-Type": "multipart/form-data"
            }
        })
    }
}

export const ApiFile = new m_api()
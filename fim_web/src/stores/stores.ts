import {defineStore} from "pinia";
import {reactive} from "vue";
import {Valid} from "@/utils/valid.ts";
import {ApiAuth, typeAuth} from "@/api/api_auth.ts";
import router from "@/router";
import {ElMessage} from "element-plus";
import {typeUser} from "@/api/api_user.ts";

interface Config {
    menuIcon: string[]
}

const config = reactive<Config>({
    menuIcon: [
        "iconfont icon-sousuo",
        "iconfont icon-QQ",
    ]
})

export const useStore = defineStore('counter', {
    state() {
        return {
            config: config,
            userInfo: typeUser.userInfo(),
        }
    },
    actions: {
        async setToken(token: string) {
            const payload = Valid.Jwt.Parse(token);
            this.userInfo = payload.PayLoad;
            this.userInfo.token = token;
            localStorage.setItem("userInfo", JSON.stringify(this.userInfo));
            router.push({name: "web"})
            ElMessage.success("登录成功")
        },
        getToken() {
            try {
                const value = localStorage.getItem("userInfo");
                if (!value) return;
                this.userInfo = JSON.parse(value)
            } catch (e) {
                localStorage.removeItem("userInfo");
            }
            console.log("userInfo", this.userInfo);
        }
    },
    getters: {},
})

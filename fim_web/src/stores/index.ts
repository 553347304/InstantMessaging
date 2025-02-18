import {defineStore} from "pinia";
import {reactive} from "vue";
import {Valid} from "@/utils/valid.ts";
import router from "@/router";
import {ElMessage} from "element-plus";
import {ApiSetting, typeSetting} from "@/api/api_setting.ts";
import {ApiUser, typeUser} from "@/api/api_user.ts";
import {typeAuth} from "@/api/api_auth.ts";
import {baseResponse} from "@/api/response.ts";

interface Config {
    menuIcon: string[]
}

const config = reactive<Config>({
    menuIcon: [
        "iconfont icon-sousuo",
        "iconfont icon-QQ",
    ]
})
const local = {
    setting: "setting",
    auth: "auth",
    userInfo: "userInfo",
}

function getLocalStorage(item: string) {
    try {
        const value = localStorage.getItem(item);
        if (value) return JSON.parse(value);
    } catch (e) {
        localStorage.removeItem(item);
    }
    return false;
}

export const useStore = defineStore('counter', {
    state() {
        return {
            config: config,
            auth: typeAuth.userInfo(),
            userInfo: typeUser.infoResponse(),
            setting: typeSetting.infoResponse(),
        }
    },
    actions: {
        async init() {
            await this.loadSetting();
            this.loadToken();
            if (!this.isLogin) return;
            await this.loadUserInfo();
        },
        async loadSetting(refresh?: boolean) {
            if (refresh) localStorage.removeItem(local.setting);
            const value = getLocalStorage(local.setting);
            if (value) {
                this.setting = value;
            } else {
                let response = await ApiSetting.Info()
                if (!baseResponse(response)) return;
                this.setting = response.data;
                localStorage.setItem(local.setting, JSON.stringify(this.setting));
            }
        },
        async loadUserInfo(refresh?: boolean) {
            if (refresh) localStorage.removeItem(local.userInfo);
            const value = getLocalStorage(local.userInfo);
            if (value) {
                this.userInfo = value;
            } else {
                let response = await ApiUser.Info()
                if (!baseResponse(response)) return;
                this.userInfo = response.data;
                localStorage.setItem(local.userInfo, JSON.stringify(this.userInfo));
            }
        },
        loadToken() {
            const value = getLocalStorage(local.auth);
            if (value) this.auth = value;
        },
        async setToken(token: string) {
            const payload = Valid.Jwt.Parse(token);
            this.auth = payload.PayLoad;
            this.auth.token = token;
            localStorage.setItem(local.auth, JSON.stringify(this.auth));
            await router.push({name: "web"});
            await this.loadUserInfo(true);
            ElMessage.success("登录成功");
        },
    },
    getters: {
        isLogin(): boolean {
            // exp的时间戳-现在的时间戳  为正就是没有过期
            return this.auth.token != ""
        },
    },
})

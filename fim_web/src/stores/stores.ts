import {defineStore} from "pinia";
import {reactive} from "vue";

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
        }
    },
    actions: {},
    getters: {},
})

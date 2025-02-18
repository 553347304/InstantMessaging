import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'

const VITE_SERVER_URL: string = "http://127.0.0.1:8080";

export default defineConfig({
    plugins: [
        vue(),
        vueDevTools(),
        AutoImport({resolvers: [ElementPlusResolver()],}),
        Components({resolvers: [ElementPlusResolver()],}),
    ],
    resolve: {
        alias: {'@': fileURLToPath(new URL('./src', import.meta.url))},
    },

    css: {
        preprocessorOptions: {
            scss: {
                additionalData: `@use "@/assets/mixin.scss" as *;`,
            },
        },
    },

    server: {
        allowedHosts: ['tcbyj.cn'], // 允许的主机
        host: "0.0.0.0",
        port: 80,
        proxy: {
            "/api/chat/ws": {
                target: VITE_SERVER_URL,
                changeOrigin: true,
                ws: true,
            },
            "/api/group/ws": {
                target: VITE_SERVER_URL,
                changeOrigin: true,
                ws: true,
            },
            "/api": {
                target: VITE_SERVER_URL,
                changeOrigin: true,
            }
        }
    }

})

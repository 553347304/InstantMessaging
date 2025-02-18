import {createApp} from 'vue'
import {createPinia} from 'pinia'

import '@/assets/base.css';
import '@/assets/theme.css';
import '@/assets/iconfont.css';


import 'element-plus/dist/index.css'
// import "element-plus/theme-chalk/el-message.css"
// import "element-plus/theme-chalk/el-message-box.css"

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')



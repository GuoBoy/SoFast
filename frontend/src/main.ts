import { createApp } from 'vue'
import App from './App.vue'
import 'element-plus/dist/index.css'

import router from './router'
import { createPinia } from 'pinia'

const app = createApp(App)

app.use(router).use(createPinia()).mount('#app')

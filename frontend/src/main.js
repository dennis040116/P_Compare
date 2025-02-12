import { createApp } from 'vue'
import './tailwind.css'
import App from './App.vue'
import { routes } from './route/routes'
import { createPinia } from 'pinia'
import { createRouter, createWebHistory } from 'vue-router'
import 'preline/dist/preline.js'
import 'preline'
import axios from './plugins/axios'

const app = createApp(App)

const router = createRouter({
  history: createWebHistory(),
  routes,
})

app.config.globalProperties.$axios = axios
app.use(router)
app.use(createPinia())
app.mount('#app')

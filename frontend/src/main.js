import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import './index.css'
import router from './router/router.js'
import VueLazyLoad from 'vue3-lazy'

const app = createApp(App)

app.use(router)
app.use(VueLazyLoad)
app.mount('#app')
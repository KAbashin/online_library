import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import './index.css'
import router from './router/router.js'
import lazyPlugin from '@/plugins/lazy'

const app = createApp(App)

app.config.errorHandler = (err, vm, info) => {
    console.error('Vue Global Error:', err, info)
}


app.use(router)
app.use(lazyPlugin)
app.mount('#app')
import axios from 'axios'

const instance = axios.create({
    baseURL:  import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
    headers: {
        'Content-Type': 'application/json'
    }
})

instance.interceptors.request.use(config => {
    console.log('[Axios Request]', config.method.toUpperCase(), config.url, config)
    const token = localStorage.getItem('token')
    if (token) {
        config.headers.Authorization = `Bearer ${token}`
    }
    return config
})

instance.interceptors.response.use(
    res => {
        console.log('[Axios Response]', res.status, res.config.url, res.data)
        return res
    },
    err => {
        if (err.response) {
            console.error('[Axios Error Response]', err.response.status, err.response.config.url, err.response.data)
            if (err.response.status === 401) {
                localStorage.removeItem('token')
                localStorage.removeItem('role')
                window.location.href = '/login'
            }
        } else {
            console.error('[Axios Error]', err.message)
        }
        return Promise.reject(err)
    }
)

export default instance
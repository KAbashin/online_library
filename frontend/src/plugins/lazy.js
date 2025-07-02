import VueLazyLoad from 'vue3-lazy'

export default {
    install(app) {
        app.use(VueLazyLoad, {
            loading: '/images/loader.gif', // нужны изображения
            error: '/images/error.png', // нужны изображения
            attempt: 1
        })
    }
}
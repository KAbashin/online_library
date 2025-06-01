import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Login from '@/views/Login.vue'
import Admin from '@/views/Admin.vue'
import NewUser from '@/views/NewUser.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        // Ленивая загрузка компонента
        component: () => import('@/views/Home.vue'),
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/Login.vue'),
    },
    {
        path: '/register',
        name: 'Register',
        component: () => import('@/views/Register.vue')
    },
    {
        path: '/adminbackdoor',
        name: 'Admin',
        component: () => import('@/views/Admin.vue'),
        meta: { minRole: 'admin' },
    },
    {
        path: '/new-user',
        name: 'NewUser',
        component: () => import('@/views/NewUser.vue'),
        meta: { minRole: 'new-user' },
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: () => import('@/views/NotFound.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

const roleHierarchy = {
    'new-user': 0,
    'user': 1,
    'admin': 2,
    'superadmin': 3,
}

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')
    const role = localStorage.getItem('role')

    const publicPaths = ['/login', '/register']

    if (!token && !publicPaths.includes(to.path)) {
        return next('/login')
    }

    // new-user должен сидеть только на своей странице
    if (role === 'new-user' && to.path !== '/new-user') {
        return next('/new-user')
    }

    // Проверка по minRole (весам)
    if (to.meta.minRole) {
        if (!token || !role) {
            return next('/login')
        }

        const userLevel = roleHierarchy[role] ?? -1
        const requiredLevel = roleHierarchy[to.meta.minRole]

        if (userLevel < requiredLevel) {
            return next('/')
        }
    }

    next()
})

export default router
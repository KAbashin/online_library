import { createRouter, createWebHistory } from 'vue-router'
import MainLayout from '@/layouts/MainLayout.vue'

const routes = [
    {
        path: '/',
        component: MainLayout,
        children: [
            {
                path: '',
                name: 'Home',
                component: () => import('@/views/Home.vue')
            },
            // Админ зона
            {
                path: 'adminbackdoor',
                name: 'Admin',
                component: () => import('@/views/Admin.vue'),
                meta: { minRole: 'admin' }
            },
            // Страница карантин для новых пользователей
            {
                path: 'new-user',
                name: 'NewUser',
                component: () => import('@/views/NewUser.vue'),
                meta: { minRole: 'new-user' }
            },
            // 1. Страница родительских категорий
            {   path: 'category',
                name: 'CategoryList',
                component: () => import('@/views/CategoryList.vue'),
                meta: { minRole: 'user' }
            },

            // 2. Страница родительской категории
            {
                path: 'category/:parentName-:parentId(\\d+)',
                name: 'ParentCategory',
                component: () => import('@/views/ParentCategory.vue'),
                props: true,
                meta: { minRole: 'user' }
            },

            // 3. Страница дочерней категории
            {
                path: 'category/:parentName-:parentId(\\d+)/:childName-:childId(\\d+)',
                name: 'ChildCategory',
                component: () => import('@/views/ChildCategory.vue'),
                props: true,
                meta: { minRole: 'user' }
            },

            // 4. Страница книги
            {
                path: 'book/:title-:id(\\d+)',
                name: 'Book',
                component: () => import('@/views/Book.vue'),
                props: true,
                meta: { minRole: 'user' }
            },

            // 5. Страница автора
            {
                path: 'author/:name-:id(\\d+)',
                name: 'Author',
                component: () => import('@/views/Author.vue'),
                props: true,
                meta: { minRole: 'user' }
            },

            // 6. Страница тега
            {
                path: ':tagName-:id(\\d+)',
                name: 'Tag',
                component: () => import('@/views/Tag.vue'),
                props: true,
                meta: { minRole: 'user' }
            },

            // 7. Профиль пользователя
            {
                path: 'profile-:id(\\d+)',
                name: 'Profile',
                component: () => import('@/views/Profile.vue'),
                props: true,
                meta: { minRole: 'user' }
            }
        ]
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/Login.vue')
    },
    {
        path: '/register',
        name: 'Register',
        component: () => import('@/views/Register.vue')
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
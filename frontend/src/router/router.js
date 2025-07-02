import { createRouter, createWebHistory } from 'vue-router'
import MainLayout from '@/layouts/MainLayout.vue'

import CategoryList from '@/views/category/CategoryList.vue'
import ParentCategory from '@/components/ParentCategories.vue'
import ChildCategory from '@/views/category/ChildCategory.vue'
import AuthorPage from '@/views/author/AuthorPage.vue'
import TagPage from '@/views/tag/TagPage.vue'
import BookPage from '@/views/book/BookPage.vue'
import ProfilePage from '@/views/profile/ProfilePage.vue'

const routes = [
    {
        path: '/',
        component: MainLayout,
        children: [
            { path: '', name: 'Home', component: () => import('@/views/Home.vue') },

            { path: 'adminbackdoor', name: 'Admin', component: () => import('@/views/admin/Admin.vue'), meta: { minRole: 'admin' } },
            { path: 'new-user', name: 'NewUser', component: () => import('@/views/auth/NewUser.vue'), meta: { minRole: 'new-user' } },

            // 👇 Категории, авторы, книги, теги и т.п. — тоже вложены
            { path: '', name: 'CategoryList', component: CategoryList, meta: { minRole: 'user' }},
            { path: 'category/:parentName-:parentId(\\d+)', name: 'ParentCategory', component: ParentCategory, props: true, meta: { minRole: 'user' } },
            { path: 'category/:parentName-:parentId(\\d+)/:childName-:childId(\\d+)', name: 'ChildCategory', component: ChildCategory, props: true, meta: { minRole: 'user' } },
            { path: 'author/:name-:id(\\d+)', component: AuthorPage, name: 'Author', props: true, meta: { minRole: 'user' } },
            { path: 'tag/:name-:id(\\d+)', component: TagPage, name: 'Tag', props: true, meta: { minRole: 'user' } },
            { path: 'book/:title-:id(\\d+)', component: BookPage, name: 'Book', props: true, meta: { minRole: 'user' } },
            { path: 'profile/:id(\\d+)', component: ProfilePage, name: 'Profile', props: true, meta: { minRole: 'user' } }
        ]
    },


    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/auth/Login.vue')
    },
    {
        path: '/register',
        name: 'Register',
        component: () => import('@/views/auth/Register.vue')
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: () => import('@/views/error/NotFound.vue')
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
import axios from './axios'  // централизованный axios с interceptors

export async function fetchCategoryBooks(categoryId) {
    return axios.get(`/api/categories/${categoryId}/books`)
}

export async function fetchCategoryInfo(categoryId) {
    return axios.get(`/api/categories/${categoryId}`)
}
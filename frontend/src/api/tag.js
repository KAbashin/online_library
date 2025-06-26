import axios from './axios'

export async function fetchTag(id) {
    return axios.get(`/api/tags/${id}`)
}

export async function fetchTagBooks(id) {
    return axios.get(`/api/books/tag/${id}`)
}
import axios from './axios'

export async function fetchAuthor(id) {
    return axios.get(`/api/authors/${id}`)
}

export async function fetchAuthorBooks(id) {
    return axios.get(`/api/books/author/${id}`)
}

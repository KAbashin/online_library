import axios from '@/utils/axios'

export async function getBook(id) {
    const response = await axios.get(`/api/books/${id}`)
    return response.data
}

export async function getBookExtras(id) {
    const response = await axios.get(`/api/books/${id}/extras`)
    return response.data
}
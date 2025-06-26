import axios from '@/plugins/axios'

export async function getUserProfile(id) {
    const { data } = await axios.get(`/api/users/${id}`)
    return data
}

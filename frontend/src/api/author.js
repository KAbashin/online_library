import axios from '@/api/axios'

/**
 * Получает список всех авторов.
 *
 * @async
 * @returns {Promise<Array>} Список авторов
 */
export async function fetchAllAuthors() {
    const { data } = await axios.get('/authors')
    return data
}

/**
 * Получает данные автора по его ID.
 *
 * @async
 * @param {number|string} id - ID автора
 * @returns {Promise<Object>} Информация об авторе
 */
export async function fetchAuthor(id) {
    const { data } = await axios.get(`/authors/${id}`)
    return data
}

/**
 * Получает список книг, написанных автором.
 *
 * @async
 * @param {number|string} id - ID автора
 * @returns {Promise<Array>} Список книг автора
 */
export async function fetchAuthorBooks(id) {
    const { data } = await axios.get(`/books/author/${id}`)
    return data
}

/**
 * Создаёт нового автора.
 *
 * @async
 * @param {Object} authorData - Данные нового автора
 * @param {string} authorData.name - Имя автора
 * @param {string} [authorData.bio] - Биография (опционально)
 * @returns {Promise<Object>} Созданный автор
 */
export async function createAuthor(authorData) {
    const { data } = await axios.post('/authors', authorData)
    return data
}

/**
 * Обновляет данные автора по ID.
 *
 * @async
 * @param {number|string} id - ID автора
 * @param {Object} updates - Обновлённые данные автора
 * @returns {Promise<Object>} Обновлённый автор
 */
export async function updateAuthor(id, updates) {
    const { data } = await axios.post(`/authors/${id}`, updates)
    return data
}

/**
 * Удаляет автора по ID (только для админов).
 *
 * @async
 * @param {number|string} id - ID автора
 * @returns {Promise<Object>} Результат удаления
 */
export async function deleteAuthor(id) {
    const { data } = await axios.post(`/authors/${id}/delete`)
    return data
}
import axios from '@/api/axios'

/**
 * Создаёт новый комментарий.
 *
 * @async
 * @param {Object} commentData - Данные комментария
 * @param {number|string} commentData.book_id - ID книги
 * @param {string} commentData.text - Текст комментария
 * @returns {Promise<Object>} Созданный комментарий
 */
export async function createComment(commentData) {
    const { data } = await axios.post('/comments', commentData)
    return data
}

/**
 * Обновляет существующий комментарий (только владелец или админ).
 *
 * @async
 * @param {number|string} id - ID комментария
 * @param {Object} updates - Новые данные
 * @param {string} updates.text - Обновлённый текст
 * @returns {Promise<Object>} Обновлённый комментарий
 */
export async function updateComment(id, updates) {
    const { data } = await axios.post(`/comments/${id}`, updates)
    return data
}

/**
 * Мягко удаляет комментарий (только владелец или админ).
 *
 * @async
 * @param {number|string} id - ID комментария
 * @returns {Promise<Object>} Результат удаления
 */
export async function deleteComment(id) {
    const { data } = await axios.post(`/comments/${id}/delete`)
    return data
}

/**
 * Получает комментарии к книге с пагинацией.
 *
 * @async
 * @param {number|string} bookId - ID книги
 * @param {Object} [params] - Параметры запроса
 * @param {number} [params.limit] - Количество записей
 * @param {number} [params.offset] - Смещение
 * @returns {Promise<Array>} Список комментариев
 */
export async function fetchCommentsByBook(bookId, params = {}) {
    const { data } = await axios.get(`/comments/book/${bookId}`, { params })
    return data
}

/**
 * Получает комментарии пользователя (только владелец или админ).
 *
 * @async
 * @param {number|string} userId - ID пользователя
 * @returns {Promise<Array>} Список комментариев пользователя
 */
export async function fetchCommentsByUser(userId) {
    const { data } = await axios.get(`/comments/user/${userId}`)
    return data
}

/**
 * Получает последние комментарии (например, для админпанели или главной страницы).
 *
 * @async
 * @returns {Promise<Array>} Список последних комментариев
 */
export async function fetchLastComments() {
    const { data } = await axios.get('/comments/last')
    return data
}

/**
 * Устанавливает статус комментария (только админ).
 *
 * @async
 * @param {number|string} id - ID комментария
 * @param {Object} statusData - Новый статус
 * @param {string} statusData.status - Новый статус (например, "approved", "rejected")
 * @returns {Promise<Object>} Обновлённый комментарий
 */
export async function setCommentStatus(id, statusData) {
    const { data } = await axios.post(`/comments/${id}/status`, statusData)
    return data
}
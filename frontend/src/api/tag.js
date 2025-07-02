import axios from '@/api/axios'

/**
 * Получает информацию о теге по его ID.
 *
 * @async
 * @param {number|string} id - Уникальный идентификатор тега
 * @returns {Promise<import('axios').AxiosResponse>} Ответ с данными тега
 *
 * @example
 * const response = await fetchTag(3);
 * console.log(response.data);
 */
export async function fetchTag(id) {
    return axios.get(`/tags/${id}`)
}

/**
 * Получает список книг, связанных с определённым тегом.
 *
 * @async
 * @param {number|string} id - Уникальный идентификатор тега
 * @returns {Promise<import('axios').AxiosResponse>} Ответ с книгами по тегу
 *
 * @example
 * const response = await fetchTagBooks(5);
 * console.log(response.data);
 */
export async function fetchTagBooks(id) {
    return axios.get(`/books/tag/${id}`)
}

/**
 * Выполняет поиск тегов по строке запроса.
 *
 * @async
 * @param {string} query - Поисковый запрос
 * @returns {Promise<import('axios').AxiosResponse>} Ответ со списком тегов
 *
 * @example
 * const response = await searchTags('history');
 */
export async function searchTags(query) {
    return axios.get(`/tags?query=${encodeURIComponent(query)}`)
}

/**
 * Создаёт новый тег.
 *
 * @async
 * @param {{ name: string }} data - Данные нового тега
 * @returns {Promise<import('axios').AxiosResponse>} Ответ с созданным тегом
 */
export async function createTag(data) {
    return axios.post('/tags', data)
}

/**
 * Обновляет существующий тег.
 *
 * Требуются права администратора.
 *
 * @async
 * @param {number|string} id - ID тега
 * @param {{ name: string }} data - Обновлённые данные
 * @returns {Promise<import('axios').AxiosResponse>} Ответ с обновлённым тегом
 */
export async function updateTag(id, data) {
    return axios.put(`/tags/${id}`, data)
}

/**
 * Удаляет тег.
 *
 * Требуются права администратора.
 *
 * @async
 * @param {number|string} id - ID тега
 * @returns {Promise<import('axios').AxiosResponse>} Ответ сервера
 */
export async function deleteTag(id) {
    return axios.post(`/tags/${id}/delete`)
}

/**
 * Получает все теги, назначенные книге.
 *
 * @async
 * @param {number|string} bookId - ID книги
 * @returns {Promise<import('axios').AxiosResponse>} Ответ со списком тегов
 */
export async function getTagsByBookId(bookId) {
    return axios.get(`/tags/book/${bookId}`)
}

/**
 * Назначает тег книге.
 *
 * Требуются права владельца книги или администратора.
 *
 * @async
 * @param {number} tagID - ID тега
 * @param {number} bookID - ID книги
 * @returns {Promise<import('axios').AxiosResponse>} Ответ сервера
 */
export async function assignTagToBook(tagID, bookID) {
    return axios.post('/tags/assign', {
        tag_id: tagID,
        book_id: bookID
    })
}

/**
 * Удаляет тег из книги.
 *
 * Требуются права владельца книги или администратора.
 *
 * @async
 * @param {number} tagID - ID тега
 * @param {number} bookID - ID книги
 * @returns {Promise<import('axios').AxiosResponse>} Ответ сервера
 */
export async function removeTagFromBook(tagID, bookID) {
    return axios.post('/tags/remove', {
        tag_id: tagID,
        book_id: bookID
    })
}
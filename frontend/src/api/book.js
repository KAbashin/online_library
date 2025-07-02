import axios from '@/api/axios'

/**
 * Получает информацию о книге по ID.
 * @param {number|string} id - ID книги
 * @returns {Promise<Object>} Данные книги
 */
export async function getBook(id) {
    const { data } = await axios.get(`/books/${id}`)
    return data
}

/**
 * Получает дополнительные данные о книге (например, оценки, количество комментариев и т.п.)
 * @param {number|string} id - ID книги
 * @returns {Promise<Object>} Дополнительная информация
 */
export async function getBookExtras(id) {
    const { data } = await axios.get(`/books/${id}/extras`)
    return data
}

/**
 * Ищет книги с параметрами.
 * @param {Object} params - query параметры поиска (title, author и т.п.)
 * @returns {Promise<Array>} Результаты поиска
 */
export async function searchBooks(params = {}) {
    const { data } = await axios.get('/books', { params })
    return data
}

/**
 * Получает книги по автору.
 * @param {number|string} authorId
 * @returns {Promise<Array>}
 */
export async function getBooksByAuthor(authorId) {
    const { data } = await axios.get(`/books/author/${authorId}`)
    return data
}

/**
 * Получает книги по тегу.
 * @param {number|string} tagId
 * @returns {Promise<Array>}
 */
export async function getBooksByTag(tagId) {
    const { data } = await axios.get(`/books/tag/${tagId}`)
    return data
}

/**
 * Проверяет наличие дубликатов книги по названию.
 * @param {string} title
 * @returns {Promise<Array>}
 */
export async function getDuplicateBooks(title) {
    const { data } = await axios.get(`/books/duplicates/${encodeURIComponent(title)}`)
    return data
}

/**
 * Получает книги текущего пользователя.
 * @returns {Promise<Array>}
 */
export async function getUserBooks() {
    const { data } = await axios.get('/books/mine')
    return data
}

/**
 * Получает новые релизы книг.
 * @returns {Promise<Array>}
 */
export async function getNewReleases() {
    const { data } = await axios.get('/books/new-releases')
    return data
}

/**
 * Получает избранные книги пользователя.
 * @returns {Promise<Array>}
 */
export async function getFavoriteBooks() {
    const { data } = await axios.get('/books/favorites')
    return data
}

/**
 * Добавляет книгу в избранное.
 * @param {number|string} bookId
 * @returns {Promise<Object>}
 */
export async function addToFavorites(bookId) {
    const { data } = await axios.post(`/books/${bookId}/favorite/add`)
    return data
}

/**
 * Удаляет книгу из избранного.
 * @param {number|string} bookId
 * @returns {Promise<Object>}
 */
export async function removeFromFavorites(bookId) {
    const { data } = await axios.post(`/books/${bookId}/favorite/remove`)
    return data
}

/**
 * Создаёт новую книгу.
 * @param {Object} bookData
 * @returns {Promise<Object>}
 */
export async function createBook(bookData) {
    const { data } = await axios.post('/books', bookData)
    return data
}

/**
 * Обновляет книгу (владелец или админ).
 * @param {number|string} bookId
 * @param {Object} updates
 * @returns {Promise<Object>}
 */
export async function updateBook(bookId, updates) {
    const { data } = await axios.post(`/books/${bookId}`, updates)
    return data
}

/**
 * Удаляет книгу (владелец или админ).
 * @param {number|string} bookId
 * @returns {Promise<Object>}
 */
export async function deleteBook(bookId) {
    const { data } = await axios.post(`/books/${bookId}/delete`)
    return data
}

/**
 * Обновляет статус книги (только админ).
 * @param {number|string} bookId
 * @param {Object} statusData
 * @returns {Promise<Object>}
 */
export async function updateBookStatus(bookId, statusData) {
    const { data } = await axios.post(`/books/${bookId}/status`, statusData)
    return data
}

/**
 * Устанавливает авторов книги (заменяет всех).
 * @param {number|string} bookId
 * @param {Array<number>} authorIds
 * @returns {Promise<Object>}
 */
export async function setBookAuthors(bookId, authorIds) {
    const { data } = await axios.post(`/books/${bookId}/authors`, { author_ids: authorIds })
    return data
}

/**
 * Добавляет автора к книге.
 * @param {number|string} bookId
 * @param {number|string} authorId
 * @returns {Promise<Object>}
 */
export async function addBookAuthor(bookId, authorId) {
    const { data } = await axios.post(`/books/${bookId}/authors/${authorId}`)
    return data
}

/**
 * Удаляет автора из книги.
 * @param {number|string} bookId
 * @param {number|string} authorId
 * @returns {Promise<Object>}
 */
export async function removeBookAuthor(bookId, authorId) {
    const { data } = await axios.post(`/books/${bookId}/authors/${authorId}/remove`)
    return data
}

/**
 * Устанавливает теги книги (заменяет все).
 * @param {number|string} bookId
 * @param {Array<number>} tagIds
 * @returns {Promise<Object>}
 */
export async function setBookTags(bookId, tagIds) {
    const { data } = await axios.post(`/books/${bookId}/tags`, { tag_ids: tagIds })
    return data
}

/**
 * Добавляет тег к книге.
 * @param {number|string} bookId
 * @param {number|string} tagId
 * @returns {Promise<Object>}
 */
export async function addBookTag(bookId, tagId) {
    const { data } = await axios.post(`/books/${bookId}/tags/${tagId}`)
    return data
}

/**
 * Удаляет тег из книги.
 * @param {number|string} bookId
 * @param {number|string} tagId
 * @returns {Promise<Object>}
 */
export async function removeBookTag(bookId, tagId) {
    const { data } = await axios.post(`/books/${bookId}/tags/${tagId}/remove`)
    return data
}
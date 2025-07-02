import axios from '@/api/axios'

/**
 * Получает профиль пользователя по ID.
 *
 * @async
 * @param {number|string} id - ID пользователя
 * @returns {Promise<Object>} Данные пользователя
 *
 * @example
 * const user = await getUserProfile(12);
 * console.log(user.name);
 */
export async function getUserProfile(id) {
    const { data } = await axios.get(`/users/${id}`)
    return data
}

/**
 * Получает список всех пользователей (только для админа).
 *
 * @async
 * @returns {Promise<Array>} Список пользователей
 *
 * @example
 * const users = await getUsers();
 * console.log(users.length);
 */
export async function getUsers() {
    const { data } = await axios.get('/users')
    return data
}

/**
 * Создаёт нового пользователя (только для админа).
 *
 * @async
 * @param {Object} userData - Данные пользователя
 * @param {string} userData.name - Имя
 * @param {string} userData.email - Email
 * @param {string} userData.password - Пароль
 * @param {string} [userData.role] - Роль (по умолчанию new-user)
 * @returns {Promise<Object>} Созданный пользователь
 */
export async function createUser(userData) {
    const { data } = await axios.post('/users', userData)
    return data
}

/**
 * Обновляет данные пользователя (владелец или админ).
 *
 * @async
 * @param {number|string} id - ID пользователя
 * @param {Object} updates - Обновлённые поля
 * @returns {Promise<Object>} Обновлённый пользователь
 */
export async function updateUser(id, updates) {
    const { data } = await axios.put(`/users/${id}`, updates)
    return data
}

/**
 * Мягкое удаление пользователя (soft delete) — только для администратора.
 *
 * @async
 * @param {number|string} id - ID пользователя
 * @returns {Promise<Object>} Результат удаления
 */
export async function softDeleteUser(id) {
    const { data } = await axios.post(`/users/${id}/delete`)
    return data
}

/**
 * Жёсткое удаление пользователя (hard delete) — только для супер-админа.
 *
 * @async
 * @param {number|string} id - ID пользователя
 * @returns {Promise<Object>} Результат удаления
 */
export async function hardDeleteUser(id) {
    const { data } = await axios.post(`/users/${id}/harddelete`)
    return data
}

/**
 * Получает книги пользователя (его загруженные книги).
 *
 * @async
 * @param {number|string} userId
 * @returns {Promise<Array>} Книги пользователя
 */
export async function getUserBooks(userId) {
    const { data } = await axios.get(`/books/author/${userId}`)
    return data
}

/**
 * Получает избранные книги пользователя.
 *
 * @async
 * @returns {Promise<Array>} Избранные книги (текущего авторизованного пользователя)
 */
export async function getUserFavoriteBooks() {
    const { data } = await axios.get('/books/favorites')
    return data
}

/**
 * Получает комментарии пользователя по ID (только если ты владелец или админ).
 *
 * @async
 * @param {number|string} userId
 * @returns {Promise<Array>} Комментарии пользователя
 */
export async function getUserComments(userId) {
    const { data } = await axios.get(`/comments/user/${userId}`)
    return data
}

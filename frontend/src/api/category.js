import axios from '@/api/axios'

/**
 * Получает все категории (в виде дерева).
 *
 * @async
 * @returns {Promise<Array>} Список всех категорий
 */
export async function fetchAllCategories() {
    const { data } = await axios.get('/categories')
    return data
}

/**
 * Получает только корневые (верхнего уровня) категории.
 *
 * @async
 * @returns {Promise<Array>} Список корневых категорий
 */
export async function fetchRootCategories() {
    const { data } = await axios.get('/categories/root')
    return data
}

/**
 * Получает информацию о категории по ID.
 *
 * @async
 * @param {number|string} categoryId - ID категории
 * @returns {Promise<Object>} Данные категории
 */
export async function fetchCategoryInfo(categoryId) {
    const { data } = await axios.get(`/categories/${categoryId}`)
    return data
}

/**
 * Получает дочерние категории для заданной категории.
 *
 * @async
 * @param {number|string} categoryId - ID родительской категории
 * @returns {Promise<Array>} Список дочерних категорий
 */
export async function fetchCategoryChildren(categoryId) {
    const { data } = await axios.get(`/categories/${categoryId}/children`)
    return data
}

/**
 * Получает список книг в категории.
 *
 * @async
 * @param {number|string} categoryId - ID категории
 * @returns {Promise<Array>} Книги в категории
 */
export async function fetchCategoryBooks(categoryId) {
    const { data } = await axios.get(`/categories/${categoryId}/books`)
    return data
}

/**
 * Создаёт новую категорию (только админ).
 *
 * @async
 * @param {Object} categoryData - Данные категории
 * @param {string} categoryData.name - Название
 * @param {number|null} [categoryData.parentId] - ID родителя (если это подкатегория)
 * @returns {Promise<Object>} Созданная категория
 */
export async function createCategory(categoryData) {
    const { data } = await axios.post('/categories', categoryData)
    return data
}

/**
 * Обновляет категорию по ID (только админ).
 *
 * @async
 * @param {number|string} categoryId - ID категории
 * @param {Object} updates - Поля для обновления
 * @returns {Promise<Object>} Обновлённая категория
 */
export async function updateCategory(categoryId, updates) {
    const { data } = await axios.post(`/categories/${categoryId}`, updates)
    return data
}

/**
 * Удаляет категорию по ID (только админ).
 *
 * @async
 * @param {number|string} categoryId - ID категории
 * @returns {Promise<Object>} Результат удаления
 */
export async function deleteCategory(categoryId) {
    const { data } = await axios.post(`/categories/${categoryId}/delete`)
    return data
}
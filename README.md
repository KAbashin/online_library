# online_library

Приложение онлайн-библиотеки.

## 1. Архитектура

- Бэкенд: Go (Gin)
- База данных: PostgreSQL
- Фронтенд: Vue.js

backend/<br>
├── cmd/ # Точка входа (main.go)<br>
├── internal/<br>
│ ├── handlers/ # HTTP-обработчики (REST)<br>
│ ├── middleware/ # Авторизация, роли<br>
│ ├── models/ # Структуры БД<br>
│ ├── repository/ # Доступ к данным<br>
│ ├── routes/ # HTTP-маршруты<br>
│ ├── service/ # Бизнес-логика<br>
│ └── pkg/ # Вспомогательные пакеты (auth, utils)<br>
├── logs/<br>
├── migrations/ # SQL-миграции<br>
├── test/<br>
└── docs/ # Swagger/OpenAPI-документация<br>
<br><br>
frontend/<br>
├── public/ # Статика<br>
└── src/<br>
├── views/ # Страницы<br>
├── store/ # Pinia/Vuex<br>
└── api.js # API-запросы<br>
<br>

---

## 2. API Endpoints (REST)

### Категории:
- `GET /api/categories` – всё дерево категорий
- `GET /api/categories/root` – корневые категории
- `GET /api/categories/{id}/children` – подкатегории
- `GET /api/categories/{id}/books` – книги в категории (пагинация)
- `POST /api/categories` – создание (админ)
- `POST /api/categories/{id}` – обновление (админ)
- `POST /api/categories/{id}/delete` – удаление (админ)

### Книги:
- `GET /api/books` – поиск / фильтрация
- `GET /api/books/{id}` – детали книги
- `GET /api/books/author/{author_id}` – по автору
- `GET /api/books/tag/{tag_id}` – по тегу
- `GET /api/books/duplicates/{title}` – поиск дубликатов
- `GET /api/books/mine` – мои книги
- `POST /api/books` – создание (авторизованный пользователь)
- `POST /api/books/{id}` – редактирование (владелец/админ)
- `POST /api/books/{id}/delete` – удаление (владелец/админ)
- `POST /api/books/{id}/status` – обновление статуса (админ)
- `POST /api/books/{book_id}/authors` – установка авторов
- `POST /api/books/{book_id}/authors/{author_id}` – добавление автора
- `POST /api/books/{book_id}/authors/{author_id}/remove` – удаление автора
- `POST /api/books/{book_id}/tags` – установка тегов
- `POST /api/books/{book_id}/tags/{tag_id}` – добавление тега
- `POST /api/books/{book_id}/tags/{tag_id}/remove` – удаление тега

#### Избранное:
- `GET /api/books/favorites` – избранные книги
- `POST /api/books/{book_id}/favorite/add` – добавить в избранное
- `POST /api/books/{book_id}/favorite/remove` – убрать из избранного

### Авторы:
- `GET /api/authors` – поиск / список
- `GET /api/authors/{id}` – подробности + книги
- `POST /api/authors` – создать
- `POST /api/authors/{id}` – редактировать
- `POST /api/authors/{id}/delete` – удалить (админ)

### Комментарии:
- `GET /api/comments/book/{book_id}` – список по книге (пагинация)
- `GET /api/comments/user/{user_id}` – список по пользователю (владелец/админ)
- `GET /api/comments/last` – последние комментарии (для админов)
- `POST /api/comments` – создать комментарий
- `POST /api/comments/{id}` – редактировать (владелец/админ)
- `POST /api/comments/{id}/delete` – удалить (владелец/админ)
- `POST /api/comments/{id}/status` – модерация комментария (админ)

### Теги:
- `GET /api/tags` – поиск
- `GET /api/tags/{id}` – по ID
- `GET /api/tags/book/{bookID}` – теги книги
- `POST /api/tags` – создать
- `PUT /api/tags/{id}` – обновить (админ)
- `POST /api/tags/{id}/delete` – удалить (админ)
- `POST /api/tags/assign` – привязать к книге (владелец/админ)
- `POST /api/tags/remove` – удалить с книги (владелец/админ)

### Аутентификация:
- `POST /api/auth/login` – вход (JWT в ответе)
- `POST /api/auth/register` – регистрация

### Пользователи:
- `GET /api/users` – список пользователей (админ)
- `POST /api/users` – создание (админ)
- `PUT /api/users/{id}` – обновление (владелец/админ)
- `POST /api/users/{id}/delete` – мягкое удаление (админ)
- `POST /api/users/{id}/harddelete` – полное удаление (только суперадмин)

---

## 3. Запуск миграций

 ```bash
migrate -path ./migrations -database "postgres://librarian:pass@localhost:5432/mydb?sslmode=disable" up
```

## 4. TODO
### Комментарии:

    Модерация (проверка комментариев со статусом pending)

    Лайки / дизлайки

    Ответы на комментарии (вложенность)

    Уведомления владельцу книги о новых комментариях

    Уведомление о новых ответах

### Рейтинг:

    Рейтинг / полезность книги (лайки, оценки)

### Инфраструктура:

    Покрытие тестами

    Логирование действий
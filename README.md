# online_library
Приложение онлайн библиотеки

1. Архитектура:
Бэкенд на Go
PostgreSQL
Фронтенд на Vue.js


backend/
├── cmd/               # Точка входа (main.go)
├── internal/          # Основная логика
│   ├── handlers/      # HTTP-обработчики (REST)
│   ├── models/        # Сущности БД (структуры Go)
│   ├── repository/    # Работа с БД (PostgreSQL)
│   ├── service/       # Бизнес-логика
│   └── pkg/           # Вспомогательные пакеты (auth, utils)
├── migrations/        # SQL-миграции (например, Goose или Liquibase)
├── pkg/               # Общие пакеты (конфиги, логгер)
└── docs/              # Swagger/OpenAPI-документация

frontend/
├── public/          # статика
└── src/
    ├── views/       # страницы
    ├── store/        # состояние (Pinia/Vuex)
    └── api.js       # вызовы к бэку


API Endpoints (REST)
Категории:
GET /api/categories – список корневых категорий.
GET /api/categories/{id}/children – подкатегории.
GET /api/categories/{id}/books – книги в категории (с пагинацией).

Книги:
GET /api/books/{id} – детали книги.
GET /api/books?search=...&tags=...&author=... – поиск/фильтрация.
POST /api/books – добавление (только для админов).

Авторы:
GET /api/authors/{id} – страница автора + его книги.
GET /api/authors?name=... – поиск авторов.

Комментарии:
GET /api/books/{id}/comments – список комментариев.
POST /api/comments – добавить (только для авторизованных).

Аутентификация:
POST /api/auth/login – вход (JWT-токен в ответе).
POST /api/auth/register – регистрация.


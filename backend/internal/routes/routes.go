package routes

import (
	"online_library/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		// Категории
		api.GET("/categories", handlers.GetRootCategories)                // список корневых категорий.
		api.GET("/categories/:id/children", handlers.GetCategoryChildren) //  подкатегории.
		api.GET("/categories/:id/books", handlers.GetBooksByCategory)     // книги в категории (с пагинацией).

		// Книги
		api.GET("/books/:id", handlers.GetBookByID) // детали книги.
		api.GET("/books", handlers.SearchBooks)     // поиск/фильтрация.
		api.POST("/books", handlers.CreateBook)     // Middleware на админа нужно отдельно // добавление (только для админов).

		// Авторы
		api.GET("/authors/:id", handlers.GetAuthorByID) //  страница автора + его книги.
		api.GET("/authors", handlers.SearchAuthors)     // поиск авторов.

		// Комментарии
		api.GET("/books/:id/comments", handlers.GetCommentsForBook) // список комментариев.
		api.POST("/comments", handlers.CreateComment)               // Middleware на auth нужно отдельно  // добавить (только для авторизованных).

		// Аутентификация
		api.POST("/auth/login", handlers.Login)       // вход (JWT-токен в ответе).
		api.POST("/auth/register", handlers.Register) // регистрация.
	}
}

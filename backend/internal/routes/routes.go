package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"online_library/backend/internal/handlers"
	"online_library/backend/internal/middleware"
	"online_library/backend/internal/repository"
	"online_library/backend/internal/service"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//bookHandler := handlers.NewBookHandler(db)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, userService)
	authHandler := handlers.NewAuthHandler(authService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	api := r.Group("/api")
	{
		// Категории
		api.GET("/categories", categoryHandler.GetAllCategories) // всё дерево категорий
		api.GET("/categories/root", categoryHandler.GetRootCategories)
		api.GET("/categories/:id", categoryHandler.GetCategoryByID)
		api.GET("/categories/:id/children", categoryHandler.GetCategoryChildren)
		api.GET("/categories/:id/books", categoryHandler.GetBooksInCategory)

		api.POST("/categories", middleware.AdminOnly(), categoryHandler.CreateCategory)
		api.POST("/categories/:id", middleware.AdminOnly(), categoryHandler.UpdateCategory)
		api.DELETE("/categories/:id", middleware.AdminOnly(), categoryHandler.DeleteCategory)

		// Книги
		/*
			api.GET("/books/:id", handlers.GetBookByID)                        // детали книги.
			api.GET("/books", handlers.SearchBooks)                            // поиск/фильтрация.
			api.POST("/books", middleware.AdminOnly(), bookHandler.CreateBook) // Middleware на админа нужно отдельно // добавление (только для админов).
		*/

		// Авторы
		api.GET("/authors/:id", handlers.GetAuthorByID) //  страница автора + его книги.
		api.GET("/authors", handlers.SearchAuthors)     // поиск авторов.

		// Комментарии
		api.GET("/books/:id/comments", handlers.GetCommentsForBook)              // список комментариев.
		api.POST("/comments", middleware.AuthRequired(), handlers.CreateComment) // Middleware на auth нужно отдельно  // добавить (только для авторизованных).

		// Пользователи
		api.GET("/users", middleware.AdminOnly(), userHandler.GetUsers)
		api.POST("/users", middleware.AdminOnly(), userHandler.СreateUser)
		api.PUT("/users/:id", middleware.AuthRequired(), userHandler.UpdateUser)
		api.DELETE("/users/:id", middleware.AdminOnly(), userHandler.SoftDeleteUser)
		api.DELETE("/users/:id/hard", middleware.SuperAdminOnly(), userHandler.HardDeleteUser)

		// Аутентификация
		api.POST("/auth/login", authHandler.Login)       // вход (JWT-токен в ответе).
		api.POST("/auth/register", authHandler.Register) // регистрация.
	}
}

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

	tagRepo := repository.NewTagRepository(db)
	tagService := service.NewTagService(tagRepo)
	tagHandler := handlers.NewTagHandler(tagService)

	// Категории
	apiCategories := r.Group("/api/categories")
	{
		apiCategories.GET("", categoryHandler.GetAllCategories) // всё дерево категорий
		apiCategories.GET("/root", categoryHandler.GetRootCategories)
		apiCategories.GET("/:id", categoryHandler.GetCategoryByID)
		apiCategories.GET("/:id/children", categoryHandler.GetCategoryChildren)
		apiCategories.GET("/:id/books", categoryHandler.GetBooksInCategory)

		apiCategories.POST("", middleware.AdminOnly(), categoryHandler.CreateCategory)
		apiCategories.POST("/:id", middleware.AdminOnly(), categoryHandler.UpdateCategory)
		apiCategories.DELETE("/:id", middleware.AdminOnly(), categoryHandler.DeleteCategory)

		authorRepo := repository.NewAuthorRepository(db)
		authorService := service.NewAuthorService(authorRepo)
		authorHandler := handlers.NewAuthorHandler(authorService)

	}

	// Книги
	/*
			api.GET("/books/:id", handlers.GetBookByID)                        // детали книги.
			api.GET("/books", handlers.SearchBooks)                            // поиск/фильтрация.
			api.POST("/books", middleware.AdminOnly(), bookHandler.CreateBook) // Middleware на админа нужно отдельно // добавление (только для админов).
		    api.GET("/books/:id/comments", handlers.GetCommentsForBook)              // список комментариев.
	*/

	// Авторы
	apiAuthors := r.Group("/api/authors")
	{
		apiAuthors.GET("", authorHandler.ListAuthors)
		apiAuthors.GET("/:id", authorHandler.GetAuthorByID)
		apiAuthors.POST("", middleware.AuthRequired(), middleware.AdminOnly(), authorHandler.CreateAuthor)
		apiAuthors.PUT("/:id", middleware.AuthRequired(), middleware.AdminOnly(), authorHandler.UpdateAuthor)
		apiAuthors.DELETE("/:id", middleware.AuthRequired(), middleware.AdminOnly(), authorHandler.DeleteAuthor)
	}

	// Комментарии
	apiComments := r.Group("/api/comments")
	{
		apiComments.POST("/comments", middleware.AuthRequired(), handlers.CreateComment) // Middleware на auth нужно отдельно  // добавить (только для авторизованных).
	}

	// Пользователи
	apiUsers := r.Group("/api/users")
	{
		apiUsers.GET("", middleware.AdminOnly(), userHandler.GetUsers)
		apiUsers.POST("", middleware.AdminOnly(), userHandler.СreateUser)
		apiUsers.PUT("/:id", middleware.AuthRequired(), userHandler.UpdateUser)
		apiUsers.DELETE("/:id", middleware.AdminOnly(), userHandler.SoftDeleteUser)
		apiUsers.DELETE("/:id/hard", middleware.SuperAdminOnly(), userHandler.HardDeleteUser)
	}

	// Аутентификация
	apiAuth := r.Group("/api/auth")
	{
		apiAuth.POST("/login", authHandler.Login)       // вход (JWT-токен в ответе).
		apiAuth.POST("/register", authHandler.Register) // регистрация.
	}

	// Теги
	apiTags := r.Group("/api/tags")
	{
		apiTags.GET("", tagHandler.SearchTags) // ?query=
		apiTags.GET("/:id", tagHandler.GetTagByID)
		apiTags.POST("", middleware.AuthRequired(), middleware.AdminOnly(), tagHandler.CreateTag)
		apiTags.PUT("/:id", middleware.AuthRequired(), middleware.AdminOnly(), tagHandler.UpdateTag)
		apiTags.DELETE("/:id", middleware.AuthRequired(), middleware.AdminOnly(), tagHandler.DeleteTag)
		apiTags.GET("/book/:bookID", tagHandler.GetTagsByBookID)
		apiTags.POST("/assign", middleware.AuthRequired(), middleware.AdminOnly(), tagHandler.AssignTagToBook)
		apiTags.DELETE("/remove", middleware.AuthRequired(), middleware.AdminOnly(), tagHandler.RemoveTagFromBook)
	}
}

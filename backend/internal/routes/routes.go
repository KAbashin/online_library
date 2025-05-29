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

	authorRepo := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepo)
	authorHandler := handlers.NewAuthorHandler(authorService)

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

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
	}

	// Книги
	// TODO проработать права доступа !!!!
	apiBooks := r.Group("/api/books")
	{
		// Публичные
		apiBooks.GET("", middleware.AuthRequired(), bookHandler.SearchBooks)
		apiBooks.GET("/:id", middleware.AuthRequired(), bookHandler.GetBookByID)
		apiBooks.GET("/author/:author_id", middleware.AuthRequired(), bookHandler.GetBooksByAuthor)
		apiBooks.GET("/tag/:tag_id", middleware.AuthRequired(), bookHandler.GetBooksByTag)
		apiBooks.GET("/duplicates/:title", middleware.AuthRequired(), bookHandler.GetDuplicateBooks)
		apiBooks.GET("/favorites", middleware.AuthRequired(), bookHandler.GetUserFavoriteBooks)
		apiBooks.GET("/mine", middleware.AuthRequired(), bookHandler.GetUserBooks)

		// Избранное
		apiBooks.POST("/:book_id/favorite", middleware.AuthRequired(), bookHandler.AddBookToFavorites)
		apiBooks.DELETE("/:book_id/favorite", middleware.AuthRequired(), bookHandler.RemoveBookFromFavorites)

		// CRUD
		apiBooks.POST("", middleware.AuthRequired(), bookHandler.CreateBook)
		apiBooks.POST("/:id", middleware.AuthRequired(), bookHandler.UpdateBook)
		apiBooks.DELETE("/:id", middleware.AuthRequired(), bookHandler.DeleteBook)

		// Статус
		apiBooks.POST("/:book_id/status", middleware.AuthRequired(), middleware.AdminOnly(), bookHandler.UpdateBookStatus)

		// Авторы
		apiBooks.POST("/:book_id/authors", middleware.AuthRequired(), bookHandler.SetBookAuthors)
		apiBooks.POST("/:book_id/authors/:author_id", middleware.AuthRequired(), bookHandler.AddBookAuthor)
		apiBooks.DELETE("/:book_id/authors/:author_id", middleware.AuthRequired(), bookHandler.RemoveBookAuthor)

		// Теги
		apiBooks.POST("/:book_id/tags", middleware.AuthRequired(), middleware.AdminOnly(), bookHandler.SetBookTags)
		apiBooks.POST("/:book_id/tags/:tag_id", middleware.AuthRequired(), middleware.AdminOnly(), bookHandler.AddBookTag)
		apiBooks.DELETE("/:book_id/tags/:tag_id", middleware.AuthRequired(), middleware.AdminOnly(), bookHandler.RemoveBookTag)
	}
	// Авторы
	apiAuthors := r.Group("/api/authors")
	{
		apiAuthors.GET("", authorHandler.ListAuthors)
		apiAuthors.GET("/:id", authorHandler.GetAuthorByID)
		apiAuthors.POST("", middleware.AuthRequired(), middleware.AdminOnly(), authorHandler.CreateAuthor)
		apiAuthors.POST("/:id", middleware.AuthRequired(), middleware.AdminOnly(), authorHandler.UpdateAuthor)
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

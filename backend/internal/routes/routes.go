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
		apiCategories.GET("", middleware.AuthRequired(), categoryHandler.GetAllCategories) // всё дерево категорий
		apiCategories.GET("/root", middleware.AuthRequired(), categoryHandler.GetRootCategories)
		apiCategories.GET("/:id", middleware.AuthRequired(), categoryHandler.GetCategoryByID)
		apiCategories.GET("/:id/children", middleware.AuthRequired(), categoryHandler.GetCategoryChildren)
		apiCategories.GET("/:id/books", middleware.AuthRequired(), categoryHandler.GetBooksInCategory)

		apiCategories.POST("", middleware.AuthRequired(), middleware.AdminOnly(), categoryHandler.CreateCategory)
		apiCategories.POST("/:id", middleware.AuthRequired(), middleware.AdminOnly(), categoryHandler.UpdateCategory)
		apiCategories.POST("/:id/delete", middleware.AuthRequired(), middleware.AdminOnly(), categoryHandler.DeleteCategory)
	}

	// Книги
	apiBooks := r.Group("/api/books")
	{
		// Публичные
		apiBooks.GET("", middleware.AuthRequired(), bookHandler.SearchBooks)
		apiBooks.GET("/:id", middleware.AuthRequired(), bookHandler.GetBookByID)
		apiBooks.GET("/author/:author_id", middleware.AuthRequired(), bookHandler.GetBooksByAuthor)
		apiBooks.GET("/tag/:tag_id", middleware.AuthRequired(), bookHandler.GetBooksByTag)
		apiBooks.GET("/duplicates/:title", middleware.AuthRequired(), bookHandler.GetDuplicateBooks)
		apiBooks.GET("/mine", middleware.AuthRequired(), bookHandler.GetUserBooks)

		// Избранное
		apiBooks.GET("/favorites", middleware.AuthRequired(), bookHandler.GetUserFavoriteBooks)
		apiBooks.POST("/:book_id/favorite/add", middleware.AuthRequired(), bookHandler.AddBookToFavorites)
		apiBooks.POST("/:book_id/favorite/remove", middleware.AuthRequired(), bookHandler.RemoveBookFromFavorites)

		// CRUD
		apiBooks.POST("", middleware.AuthRequired(), bookHandler.CreateBook)
		apiBooks.POST("/:id", middleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.UpdateBook)
		apiBooks.POST("/:id/delete", middleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.DeleteBook)

		// Статус
		apiBooks.POST("/:book_id/status", middleware.AuthRequired(), middleware.AdminOnly(), bookHandler.UpdateBookStatus)

		// Авторы
		apiBooks.POST("/:book_id/authors", middleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.SetBookAuthors)
		apiBooks.POST("/:book_id/authors/:author_id", middleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.AddBookAuthor)
		apiBooks.POST("/:book_id/authors/:author_id/remove", middleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.RemoveBookAuthor)

		// Теги
		apiBooks.POST("/:book_id/tags", middleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.SetBookTags)
		apiBooks.POST("/:book_id/tags/:tag_id", middleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.AddBookTag)
		apiBooks.POST("/:book_id/tags/:tag_id/remove", middleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.RemoveBookTag)
	}
	// Авторы
	apiAuthors := r.Group("/api/authors")
	{
		apiAuthors.GET("", middleware.AuthRequired(), authorHandler.ListAuthors)
		apiAuthors.GET("/:id", middleware.AuthRequired(), authorHandler.GetAuthorByID)
		apiAuthors.POST("", middleware.AuthRequired(), authorHandler.CreateAuthor)
		apiAuthors.POST("/:id", middleware.AuthRequired(), authorHandler.UpdateAuthor)
		apiAuthors.POST("/:id/delete", middleware.AuthRequired(), middleware.AdminOnly(), authorHandler.DeleteAuthor)
	}

	// Комментарии
	apiComments := r.Group("/api/comments")
	{
		apiComments.POST("/comments", middleware.AuthRequired(), handlers.CreateComment) // Middleware на auth нужно отдельно  // добавить (только для авторизованных).
	}

	// Пользователи
	apiUsers := r.Group("/api/users")
	{
		apiUsers.GET("", middleware.AuthRequired(), middleware.AdminOnly(), userHandler.GetUsers)
		apiUsers.POST("", middleware.AuthRequired(), middleware.AdminOnly(), userHandler.СreateUser)
		apiUsers.PUT("/:id", middleware.AuthRequired(), middleware.OwnerOrAdmin(), userHandler.UpdateUser)
		apiUsers.POST("/:id/delete", middleware.AuthRequired(), middleware.AdminOnly(), userHandler.SoftDeleteUser)
		apiUsers.POST("/:id/harddelete", middleware.AuthRequired(), middleware.SuperAdminOnly(), userHandler.HardDeleteUser)
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
		apiTags.GET("", middleware.AuthRequired(), tagHandler.SearchTags) // ?query=
		apiTags.GET("/:id", middleware.AuthRequired(), tagHandler.GetTagByID)
		apiTags.POST("", middleware.AuthRequired(), tagHandler.CreateTag)
		apiTags.PUT("/:id", middleware.AuthRequired(), middleware.AdminOnly(), tagHandler.UpdateTag)
		apiTags.POST("/:id/delete", middleware.AuthRequired(), middleware.AdminOnly(), tagHandler.DeleteTag)
		apiTags.GET("/book/:bookID", middleware.AuthRequired(), tagHandler.GetTagsByBookID)
		apiTags.POST("/assign", middleware.AuthRequired(), middleware.OwnerOrAdmin(), tagHandler.AssignTagToBook)
		apiTags.POST("/remove", middleware.AuthRequired(), middleware.OwnerOrAdmin(), tagHandler.RemoveTagFromBook)
	}
}

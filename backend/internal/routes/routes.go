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

	authMiddleware := middleware.NewAuthMiddleware(userService)

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

	commentRepo := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := handlers.NewCommentHandler(commentService)

	bookRepo := repository.NewBookRepository(db)
	imageRepo := repository.NewImageRepository(db)
	fileRepo := repository.NewFileRepository(db)
	bookService := service.NewBookService(bookRepo, tagRepo, imageRepo, fileRepo, commentRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	// Категории
	apiCategories := r.Group("/api/categories")
	{
		apiCategories.GET("", authMiddleware.AuthRequired(), categoryHandler.GetAllCategories) // всё дерево категорий
		apiCategories.GET("/root", authMiddleware.AuthRequired(), categoryHandler.GetRootCategories)
		apiCategories.GET("/:id", authMiddleware.AuthRequired(), categoryHandler.GetCategoryByID)
		apiCategories.GET("/:id/children", authMiddleware.AuthRequired(), categoryHandler.GetCategoryChildren)
		apiCategories.GET("/:id/books", authMiddleware.AuthRequired(), categoryHandler.GetBooksInCategory)

		apiCategories.POST("", authMiddleware.AuthRequired(), middleware.AdminOnly(), categoryHandler.CreateCategory)
		apiCategories.POST("/:id", authMiddleware.AuthRequired(), middleware.AdminOnly(), categoryHandler.UpdateCategory)
		apiCategories.POST("/:id/delete", authMiddleware.AuthRequired(), middleware.AdminOnly(), categoryHandler.DeleteCategory)
	}

	// Книги
	apiBooks := r.Group("/api/books")
	{
		//
		apiBooks.GET("", authMiddleware.AuthRequired(), bookHandler.SearchBooks)
		apiBooks.GET("/:book_id", authMiddleware.AuthRequired(), bookHandler.GetBookByID)
		apiBooks.GET("/:book_id/extras", authMiddleware.AuthRequired(), bookHandler.GetBookExtras)
		apiBooks.GET("/author/:author_id", authMiddleware.AuthRequired(), bookHandler.GetBooksByAuthor)
		apiBooks.GET("/tag/:tag_id", authMiddleware.AuthRequired(), bookHandler.GetBooksByTag)
		apiBooks.GET("/duplicates/:title", authMiddleware.AuthRequired(), bookHandler.GetDuplicateBooks)
		apiBooks.GET("/mine", authMiddleware.AuthRequired(), bookHandler.GetUserBooks)
		apiBooks.GET("/new-releases", authMiddleware.AuthRequired(), bookHandler.GetNewReleases)

		// Избранное
		apiBooks.GET("/favorites", authMiddleware.AuthRequired(), bookHandler.GetUserFavoriteBooks)
		apiBooks.POST("/:book_id/favorite/add", authMiddleware.AuthRequired(), bookHandler.AddBookToFavorites)
		apiBooks.POST("/:book_id/favorite/remove", authMiddleware.AuthRequired(), bookHandler.RemoveBookFromFavorites)

		// CRUD
		apiBooks.POST("", authMiddleware.AuthRequired(), bookHandler.CreateBook)
		apiBooks.POST("/:book_id", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.UpdateBook)
		apiBooks.POST("/:book_id/delete", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.DeleteBook)

		// Статус
		apiBooks.POST("/:book_id/status", authMiddleware.AuthRequired(), middleware.AdminOnly(), bookHandler.UpdateBookStatus)

		// Авторы
		apiBooks.POST("/:book_id/authors", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.SetBookAuthors)
		apiBooks.POST("/:book_id/authors/:author_id", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.AddBookAuthor)
		apiBooks.POST("/:book_id/authors/:author_id/remove", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.RemoveBookAuthor)

		// Теги
		apiBooks.POST("/:book_id/tags", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.SetBookTags)
		apiBooks.POST("/:book_id/tags/:tag_id", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.AddBookTag)
		apiBooks.POST("/:book_id/tags/:tag_id/remove", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), bookHandler.RemoveBookTag)
	}
	// Авторы
	apiAuthors := r.Group("/api/authors", authMiddleware.AuthRequired())
	{
		apiAuthors.GET("", authMiddleware.AuthRequired(), authorHandler.ListAuthors)
		apiAuthors.GET("/:id", authMiddleware.AuthRequired(), authorHandler.GetAuthorByID)
		apiAuthors.POST("", authMiddleware.AuthRequired(), authorHandler.CreateAuthor)
		apiAuthors.POST("/:id", authMiddleware.AuthRequired(), authorHandler.UpdateAuthor)
		apiAuthors.POST("/:id/delete", authMiddleware.AuthRequired(), middleware.AdminOnly(), authorHandler.DeleteAuthor)
	}

	// Комментарии
	apiComments := r.Group("/api/comments", authMiddleware.AuthRequired())
	{
		apiComments.POST("", authMiddleware.AuthRequired(), commentHandler.CreateComment)                                       // создание
		apiComments.POST("/:id", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), commentHandler.UpdateComment)        // обновление текста (автор или админ)
		apiComments.POST("/:id/delete", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), commentHandler.DeleteComment) // мягкое удаление

		apiComments.GET("/book/:book_id", authMiddleware.AuthRequired(), commentHandler.GetCommentsByBook) // пагинация ?limit=&offset=
		apiComments.GET("/user/:user_id", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), commentHandler.GetCommentsByUser)
		apiComments.GET("/last", authMiddleware.AuthRequired(), commentHandler.GetLastComments)

		apiComments.POST("/:id/status", authMiddleware.AuthRequired(), middleware.AdminOnly(), commentHandler.SetStatus)
	}

	// Пользователи
	apiUsers := r.Group("/api/users")
	{
		apiUsers.GET("", authMiddleware.AuthRequired(), middleware.AdminOnly(), userHandler.GetUsers)
		apiUsers.POST("", authMiddleware.AuthRequired(), middleware.AdminOnly(), userHandler.CreateUser)
		apiUsers.PUT("/:id", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), userHandler.UpdateUser)
		apiUsers.POST("/:id/delete", authMiddleware.AuthRequired(), middleware.AdminOnly(), userHandler.SoftDeleteUser)
		apiUsers.POST("/:id/harddelete", authMiddleware.AuthRequired(), middleware.SuperAdminOnly(), userHandler.HardDeleteUser)
	}

	// Аутентификация
	apiAuth := r.Group("/api/auth")
	{
		apiAuth.GET("/me", authMiddleware.AuthRequired(), authHandler.Me)
		apiAuth.POST("/register", authHandler.Register)
		apiAuth.POST("/login", authHandler.Login)
		apiAuth.POST("/logout", authMiddleware.AuthRequired(), authHandler.Logout)
	}

	// Теги
	apiTags := r.Group("/api/tags")
	{
		apiTags.GET("", authMiddleware.AuthRequired(), tagHandler.SearchTags) // ?query=
		apiTags.GET("/:id", authMiddleware.AuthRequired(), tagHandler.GetTagByID)
		apiTags.POST("", authMiddleware.AuthRequired(), tagHandler.CreateTag)
		apiTags.PUT("/:id", authMiddleware.AuthRequired(), middleware.AdminOnly(), tagHandler.UpdateTag)
		apiTags.POST("/:id/delete", authMiddleware.AuthRequired(), middleware.AdminOnly(), tagHandler.DeleteTag)
		apiTags.GET("/book/:bookID", authMiddleware.AuthRequired(), tagHandler.GetTagsByBookID)
		apiTags.POST("/assign", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), tagHandler.AssignTagToBook)
		apiTags.POST("/remove", authMiddleware.AuthRequired(), middleware.OwnerOrAdmin(), tagHandler.RemoveTagFromBook)
	}
}

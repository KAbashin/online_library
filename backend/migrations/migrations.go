package migrations

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
)

func RunMigrations(db *sql.DB) error {
	fmt.Println("Running SQLite DB migrations...")
	statements := []string{
		//-- Категории (древовидная структура)
		`CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		parent_id INT REFERENCES categories(id) ON DELETE CASCADE,
		slug VARCHAR(255) UNIQUE  -- Для ЧПУ-URL
	);`,

		//-- Авторы
		`CREATE TABLE IF NOT EXISTS authors (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		bio TEXT,
		photo_url VARCHAR(512)
	);`,

		//-- Книги
		`CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,          -- Аннотация
		publish_year INT,          -- Год издания
		pages INT,                 -- Количество страниц
		language VARCHAR(100),     -- Язык
		publisher VARCHAR(255),    -- Издательство
		type VARCHAR(50),          -- Тип: book/journal/article
		rating INT DEFAULT 0,      -- Рейтинг 0-100
		cover_url VARCHAR(512),    -- Обложка
		created_at TIMESTAMP DEFAULT NOW()
	);`,

		//-- Связь книг и категорий (многие-ко-многим)
		`CREATE TABLE IF NOT EXISTS book_categories (
		book_id INT REFERENCES books(id) ON DELETE CASCADE,
		category_id INT REFERENCES categories(id) ON DELETE CASCADE,
		PRIMARY KEY (book_id, category_id)
	);`,

		//-- Связь книг и авторов (многие-ко-многим)
		`CREATE TABLE IF NOT EXISTS book_authors (
		book_id INT REFERENCES books(id) ON DELETE CASCADE,
		author_id INT REFERENCES authors(id) ON DELETE CASCADE,
		PRIMARY KEY (book_id, author_id)
	);`,

		//-- Изображения книги (доп. картинки)
		`CREATE TABLE IF NOT EXISTS book_images (
		id SERIAL PRIMARY KEY,
		book_id INT REFERENCES books(id) ON DELETE CASCADE,
		url VARCHAR(512) NOT NULL,
		order_index INT  -- Порядок отображения
	);`,

		//-- Хештеги (теги)
		`CREATE TABLE IF NOT EXISTS tags (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE
	);`,

		//-- Связь книг и тегов
		`CREATE TABLE IF NOT EXISTS book_tags (
		book_id INT REFERENCES books(id) ON DELETE CASCADE,
		tag_id INT REFERENCES tags(id) ON DELETE CASCADE,
		PRIMARY KEY (book_id, tag_id)
	);`,

		//-- Комментарии/цитаты
		`CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
		book_id INT REFERENCES books(id) ON DELETE CASCADE,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		text TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`,

		//-- Пользователи (для авторизации)
		`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,  -- bcrypt
		name VARCHAR(100),
		role VARCHAR(50) DEFAULT 'user'       -- user/admin
	);`,
	}
	for _, stmt := range statements {
		tableName := extractTableName(stmt)
		fmt.Printf("Migrating table: %s\n", tableName)

		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("migration failed for table %s: %w", tableName, err)
		}
	}

	fmt.Println("DB migrations completed")
	return nil
}

func extractTableName(stmt string) string {
	re := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+IF\s+NOT\s+EXISTS\s+(\w+)`)
	matches := re.FindStringSubmatch(stmt)
	if len(matches) >= 2 {
		return strings.TrimSpace(matches[1])
	}
	return "unknown"
}

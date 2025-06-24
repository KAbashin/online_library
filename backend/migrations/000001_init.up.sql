-- Категории
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INT REFERENCES categories(id) ON DELETE CASCADE,
    slug VARCHAR(255) UNIQUE,
    description TEXT
    );

-- Авторы
CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    name_ru VARCHAR(255) NOT NULL,
    name_en VARCHAR(255) NOT NULL,
    UNIQUE(name_ru, name_en)
    );

-- Книги
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    publish_year INT,
    pages INT,
    language VARCHAR(100),
    publisher VARCHAR(255),
    type VARCHAR(50) NOT NULL DEFAULT 'book'
        CHECK (type IN ('book', 'journal', 'article', 'other')),
    rating INT DEFAULT 0,
    cover_url VARCHAR(512),
    status VARCHAR(20) NOT NULL DEFAULT 'quarantine',
    created_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
    );

-- Связь книг и категорий
CREATE TABLE IF NOT EXISTS book_categories (
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, category_id)
    );

-- Связь книг и авторов
CREATE TABLE IF NOT EXISTS book_authors (
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    author_id INT REFERENCES authors(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, author_id)
    );

-- Изображения книги
CREATE TABLE IF NOT EXISTS book_images (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    url VARCHAR(512) NOT NULL,
    order_index INT
    );

-- Хештеги
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    color VARCHAR(7) -- hex-код, напр. "#FF8800"
    );

-- Связь книг и тегов
CREATE TABLE IF NOT EXISTS book_tags (
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    tag_id INT REFERENCES tags(id) ON DELETE CASCADE,
    weight INT DEFAULT 1, -- 1 = основной, 0 = второстепенный
    PRIMARY KEY (book_id, tag_id)
    );

-- Пользователи
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'new-user',
    bio TEXT,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    token_version INT NOT NULL DEFAULT 1,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
    );

-- Комментарии
CREATE TABLE IF NOT EXISTS comments (
   id SERIAL PRIMARY KEY,
   book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
   user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
   text TEXT NOT NULL,
   created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
   updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
   status VARCHAR(20) DEFAULT 'active'
);

CREATE INDEX idx_comments_book_id ON comments(book_id);
CREATE INDEX idx_comments_user_id ON comments(user_id);
CREATE INDEX idx_comments_status ON comments(status);

-- Файлы книги
CREATE TABLE IF NOT EXISTS book_files (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    format VARCHAR(50) NOT NULL,
    url TEXT NOT NULL,
    file_size BIGINT,
    hash VARCHAR(128) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description TEXT,
    UNIQUE(book_id, format)
    );

-- Избранные книги пользователей
CREATE TABLE IF NOT EXISTS book_favorites (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, book_id)
    );
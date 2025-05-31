-- Электронные файлы книги
CREATE TABLE IF NOT EXISTS book_files (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    format VARCHAR(20) NOT NULL,  -- PDF, EPUB, MOBI и т.п.
    file_url VARCHAR(512) NOT NULL,
    size_mb DECIMAL(10, 2),
    uploaded_at TIMESTAMP DEFAULT NOW()
    );

-- Связи между книгами (переводы, издания)
CREATE TABLE IF NOT EXISTS book_links (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id) ON DELETE CASCADE,
    related_book_id INT REFERENCES books(id) ON DELETE CASCADE,
    relation_type VARCHAR(50) NOT NULL  -- "translation", "edition", "abridged", и т.д.
    );
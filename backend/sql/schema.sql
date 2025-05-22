CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    publish_year INT,
    pages INT,
    language VARCHAR(100),
    publisher VARCHAR(255),
    type VARCHAR(50),
    rating INT DEFAULT 0,
    cover_url VARCHAR(512),
    created_at TIMESTAMP DEFAULT NOW()
    );

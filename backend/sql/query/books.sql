-- name: GetBookByID :one
SELECT * FROM books WHERE id = $1;

-- name: ListBooks :many
SELECT * FROM books ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: SearchBooks :many
SELECT * FROM books
WHERE title ILIKE '%' || $1 || '%'
ORDER BY created_at DESC
    LIMIT $2 OFFSET $3;

-- name: CreateBook :one
INSERT INTO books (title, description, publish_year, pages, language, publisher, type, cover_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1;
package models

import "time"

type BookFile struct {
	ID          int       `json:"id"`
	BookID      int       `json:"book_id"`
	Format      string    `json:"format"` // Например: "pdf", "epub"
	URL         string    `json:"url"`
	FileSize    int64     `json:"file_size"` // В байтах
	Hash        string    `json:"hash"`      // SHA256 или MD5
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"` // Описание/комментарий к файлу
}

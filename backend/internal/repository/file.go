package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"online_library/backend/internal/models"
)

// FileRepository определяет методы для работы с файлами, связанными с книгами.
type FileRepository interface {
	// GetFileByID возвращает файл по его уникальному идентификатору.
	GetFileByID(id int) (models.BookFile, error)

	// CreateFile сохраняет новый файл и возвращает его ID и дату создания.
	CreateFile(tag *models.BookFile) error

	// UpdateFile обновляет данные о файле.
	UpdateFile(tag *models.BookFile) error

	// DeleteFile удаляет файл по его ID.
	DeleteFile(id int) error

	// DeleteFilesByBookID удаляет все файлы, связанные с указанной книгой.
	// Используется при удалении книги или очистке связанных ресурсов.
	DeleteFilesByBookID(bookID int) error

	// GetFilesByBookID возвращает все файлы, прикреплённые к книге, отсортированные по дате.
	GetFilesByBookID(bookID int) ([]models.BookFile, error)

	// FindFileByHash ищет файл по уникальному хэшу (SHA256 или MD5).
	// Используется для проверки наличия дубликатов.
	// Возвращает пустую структуру, если файл не найден.
	FindFileByHash(hash string) (models.BookFile, error)

	// ListRecentFiles возвращает список последних загруженных файлов,
	// отсортированных по дате создания (от новых к старым).
	// Используется в админ-панели или для вывода новинок.
	// limit — максимальное количество результатов.
	ListRecentFiles(limit int) ([]models.BookFile, error)
}

type fileRepo struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) FileRepository {
	return &fileRepo{db: db}
}

func (f *fileRepo) GetFileByID(id int) (models.BookFile, error) {
	query := `SELECT id, book_id, format, url, file_size, hash, created_at, description FROM book_files WHERE id = $1`
	var file models.BookFile
	err := f.db.QueryRow(query, id).Scan(
		&file.ID, &file.BookID, &file.Format, &file.URL, &file.FileSize,
		&file.Hash, &file.CreatedAt, &file.Description,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return file, nil
		}
		return file, fmt.Errorf("GetFileByID: %w", err)
	}
	return file, nil
}

func (f *fileRepo) CreateFile(file *models.BookFile) error {
	query := `INSERT INTO book_files (book_id, format, url, file_size, hash, description)
	          VALUES ($1, $2, $3, $4, $5, $6)
	          RETURNING id, created_at`
	err := f.db.QueryRow(query,
		file.BookID, file.Format, file.URL,
		file.FileSize, file.Hash, file.Description,
	).Scan(&file.ID, &file.CreatedAt)
	if err != nil {
		return fmt.Errorf("CreateFile: %w", err)
	}
	return nil
}

func (f *fileRepo) UpdateFile(file *models.BookFile) error {
	query := `UPDATE book_files
	          SET book_id = $1, format = $2, url = $3, file_size = $4, hash = $5, description = $6
	          WHERE id = $7`
	_, err := f.db.Exec(query,
		file.BookID, file.Format, file.URL,
		file.FileSize, file.Hash, file.Description,
		file.ID,
	)
	if err != nil {
		return fmt.Errorf("UpdateFile: %w", err)
	}
	return nil
}

func (f *fileRepo) DeleteFile(id int) error {
	query := `DELETE FROM book_files WHERE id = $1`
	_, err := f.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DeleteFile: %w", err)
	}
	return nil
}

func (f *fileRepo) GetFilesByBookID(bookID int) ([]models.BookFile, error) {
	query := `SELECT id, book_id, format, url, file_size, hash, created_at, description
	          FROM book_files WHERE book_id = $1 ORDER BY created_at DESC`
	rows, err := f.db.Query(query, bookID)
	if err != nil {
		return nil, fmt.Errorf("GetFilesByBookID: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var files []models.BookFile
	for rows.Next() {
		var file models.BookFile
		err := rows.Scan(
			&file.ID, &file.BookID, &file.Format, &file.URL,
			&file.FileSize, &file.Hash, &file.CreatedAt, &file.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("GetFilesByBookID: scan error: %w", err)
		}
		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetFilesByBookID: rows error: %w", err)
	}
	return files, nil
}

// FindFileByHash ищет файл по уникальному хэшу (SHA256 или MD5).
// Используется для проверки наличия дубликатов.
// Возвращает пустую структуру, если файл не найден.
func (f *fileRepo) FindFileByHash(hash string) (models.BookFile, error) {
	query := `SELECT id, book_id, format, url, file_size, hash, created_at, description
	          FROM book_files WHERE hash = $1 LIMIT 1`
	var file models.BookFile
	err := f.db.QueryRow(query, hash).Scan(
		&file.ID, &file.BookID, &file.Format, &file.URL,
		&file.FileSize, &file.Hash, &file.CreatedAt, &file.Description,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return file, nil
		}
		return file, fmt.Errorf("FindFileByHash: %w", err)
	}
	return file, nil
}

// DeleteFilesByBookID удаляет все файлы, связанные с указанной книгой.
// Используется при удалении книги или очистке связанных ресурсов.
func (f *fileRepo) DeleteFilesByBookID(bookID int) error {
	query := `DELETE FROM book_files WHERE book_id = $1`
	_, err := f.db.Exec(query, bookID)
	if err != nil {
		return fmt.Errorf("DeleteFilesByBookID: %w", err)
	}
	return nil
}

// ListRecentFiles возвращает список последних загруженных файлов,
// отсортированных по дате создания (от новых к старым).
// Используется в админ-панели или для вывода новинок.
// limit — максимальное количество результатов.
func (f *fileRepo) ListRecentFiles(limit int) ([]models.BookFile, error) {
	query := `SELECT id, book_id, format, url, file_size, hash, created_at, description
	          FROM book_files ORDER BY created_at DESC LIMIT $1`
	rows, err := f.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("ListRecentFiles: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var files []models.BookFile
	for rows.Next() {
		var file models.BookFile
		err := rows.Scan(
			&file.ID, &file.BookID, &file.Format, &file.URL,
			&file.FileSize, &file.Hash, &file.CreatedAt, &file.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("ListRecentFiles: scan error: %w", err)
		}
		files = append(files, file)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ListRecentFiles: rows error: %w", err)
	}
	return files, nil
}

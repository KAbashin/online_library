package dto

import (
	"fmt"
	"online_library/backend/internal/models"
	"sort"
	"time"
)

// Загружается в первую очередь, будет кешироваться
type BookDTO struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description *string        `json:"description"`
	PublishYear *int           `json:"publish_year,omitempty"`
	Pages       *int           `json:"pages,omitempty"`
	Language    *string        `json:"language,omitempty"`
	Publisher   *string        `json:"publisher,omitempty"`
	Type        *string        `json:"type,omitempty"`
	CoverURL    *string        `json:"cover_url,omitempty"`
	Authors     []AuthorDTO    `json:"authors"`
	Tags        []BookTagDTO   `json:"tags"`
	Files       []BookFileDTO  `json:"files"`
	Images      []BookImageDTO `json:"images"`
}

// Загружается во вторую очередь , не будет кешироваться
type BookExtrasDTO struct {
	// Rating      int               `json:"rating"`
	InFavorites bool             `json:"in_favorites"`
	Comments    []BookCommentDTO `json:"comments"`
	// Relations   []BookRelationDTO `json:"relations"`
}

type BookPreviewDTO struct {
	ID            int         `json:"id"`
	Title         string      `json:"title"`
	CoverURL      *string     `json:"cover_url,omitempty"`
	Authors       []AuthorDTO `json:"authors"`
	CoverImageURL *string     `json:"cover_image_url,omitempty"`
	PublishYear   *int        `json:"publish_year,omitempty"`
}

type BookFileDTO struct {
	ID          int    `json:"id"`
	Format      string `json:"format"` // Например: "pdf", "epub"
	URL         string `json:"url"`
	FileSize    int64  `json:"file_size"`   // В байтах
	Description string `json:"description"` // Описание/комментарий к файлу
}

type BookImageDTO struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	OrderIndex int    `json:"order_index"`
}

type BookRelationDTO struct {
	ID        int    `json:"id"`         // Уникальный ID связи
	RelatedID int    `json:"related_id"` // Книга, с которой есть связь
	Relation  string `json:"relation"`   // Тип связи
}

type BookTagDTO struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color,omitempty"`
	Weight int    `json:"weight"`
}

type BookCommentDTO struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` // Добавляем для отслеживания изменений
}

func ConvertBookFile(f models.BookFile) BookFileDTO {
	return BookFileDTO{
		ID:          f.ID,
		Format:      f.Format,
		URL:         f.URL,
		FileSize:    f.FileSize,
		Description: f.Description,
	}
}

func ConvertBookFiles(files []models.BookFile) []BookFileDTO {
	result := make([]BookFileDTO, 0, len(files))
	for _, f := range files {
		result = append(result, ConvertBookFile(f))
	}
	return result
}

func ConvertBookImage(img models.BookImage) BookImageDTO {
	return BookImageDTO{
		ID:         img.ID,
		URL:        img.URL,
		OrderIndex: img.OrderIndex,
	}
}

func ConvertBookImages(images []models.BookImage) []BookImageDTO {
	result := make([]BookImageDTO, 0, len(images))
	for _, img := range images {
		result = append(result, ConvertBookImage(img))
	}
	return result
}

func ConvertBookRelation(rel models.BookRelation) BookRelationDTO {
	return BookRelationDTO{
		ID:        rel.ID,
		RelatedID: rel.RelatedID,
		Relation:  rel.Relation,
	}
}

func ConvertBookRelations(rels []models.BookRelation) []BookRelationDTO {
	result := make([]BookRelationDTO, 0, len(rels))
	for _, r := range rels {
		result = append(result, ConvertBookRelation(r))
	}
	return result
}

func MergeTagsWithWeight(tags []models.Tag, bookTags []models.BookTag) []BookTagDTO {
	weights := make(map[int]int)
	for _, bt := range bookTags {
		weights[bt.TagID] = bt.Weight
	}

	var result []BookTagDTO
	for _, tag := range tags {
		result = append(result, BookTagDTO{
			ID:     tag.ID,
			Name:   tag.Name,
			Color:  tag.Color,
			Weight: weights[tag.ID],
		})
	}

	// (по желанию) сортировка по weight
	sort.Slice(result, func(i, j int) bool {
		return result[i].Weight > result[j].Weight
	})

	return result
}

func ConvertBookComment(c models.BookComment) BookCommentDTO {
	return BookCommentDTO{
		ID:        c.ID,
		UserID:    c.UserID,
		Text:      c.Text,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ConvertBookComments(comments []models.BookComment) []BookCommentDTO {
	result := make([]BookCommentDTO, 0, len(comments))
	for _, c := range comments {
		result = append(result, ConvertBookComment(c))
	}
	return result
}

func ConvertBookToPreviewDTO(book models.Book, authors []models.Author, images []models.BookImage) BookPreviewDTO {
	var coverURL *string
	for _, img := range images {
		if img.OrderIndex == 1 {
			coverURL = &img.URL
			break
		}
	}

	return BookPreviewDTO{
		ID:            book.ID,
		Title:         book.Title,
		CoverImageURL: coverURL,
		Authors:       ConvertAuthors(authors),
		PublishYear:   book.PublishYear,
	}
}

func ConvertBooksToPreviewDTOs(
	books []*models.Book,
	getAuthors func(bookID int) ([]models.Author, error),
	getImages func(bookID int) ([]models.BookImage, error),
) ([]BookPreviewDTO, error) {
	var result []BookPreviewDTO

	for _, book := range books {
		authors, err := getAuthors(book.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get authors for book %d: %w", book.ID, err)
		}

		images, err := getImages(book.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get images for book %d: %w", book.ID, err)
		}

		dtoItem := ConvertBookToPreviewDTO(*book, authors, images)
		result = append(result, dtoItem)
	}

	return result, nil
}

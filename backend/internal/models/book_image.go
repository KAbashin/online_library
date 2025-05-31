package models

type BookImage struct {
	ID         int    `json:"id"`
	BookID     int    `json:"book_id"`
	URL        string `json:"url"`
	OrderIndex int    `json:"order_index"`
}

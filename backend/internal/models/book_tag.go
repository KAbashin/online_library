package models

type BookTag struct {
	BookID int `json:"book_id"`
	TagID  int `json:"tag_id"`
	Weight int `json:"weight"`
}

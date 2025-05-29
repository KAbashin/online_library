package models

type BookRelation struct {
	ID        int    `json:"id"`         // Уникальный ID связи
	BookID    int    `json:"book_id"`    // Исходная книга
	RelatedID int    `json:"related_id"` // Книга, с которой есть связь
	Relation  string `json:"relation"`   // Тип связи, например "translation" или "reissue"
}

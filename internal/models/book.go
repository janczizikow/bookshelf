package models

import "time"

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Price       int    `json:"price"`
	Publisher   string `json:"publisher"`
	Pages       int    `json:"pages"`
	PublishDate string `json:"publish_date"`
	AuthorID    int    `json:"author_id"`

	// Timestamps

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

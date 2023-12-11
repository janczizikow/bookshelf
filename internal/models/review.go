package models

import "time"

type Review struct {
	ID          int    `json:"id"`
	Rating      int8   `json:"rating"`
	Description string `json:"description"`
	BookID      int    `json:"book_id"`
	UserID      int    `json:"user_id"`

	// Timestamps

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package models

import "time"

type OrderItem struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	BookID   int `json:"book_id"`
	Quantity int `json:"quantity"`

	// Timestamps

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

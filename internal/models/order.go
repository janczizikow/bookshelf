package models

import "time"

type Order struct {
	ID     int         `json:"id"`
	UserID int         `json:"user_id"`
	Total  int         `json:"total"`
	Items  []OrderItem `json:"items"`

	// Timestamps

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

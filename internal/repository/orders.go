package repository

import (
	"database/sql"

	"github.com/janczizikow/bookshelf/internal/models"
)

type ordersRepository struct {
	db *sql.DB
}

type OrdersRepository interface {
	Find(id int) (*models.Order, error)
}

// NewOrdersRepository returns a new instance of a orders repository.
func NewOrdersRepository(db *sql.DB) *ordersRepository {
	return &ordersRepository{db: db}
}

func (r *ordersRepository) Find(id int) (*models.Order, error) {
	order := models.Order{}
	row := r.db.QueryRow("SELECT * FROM orders WHERE id = ? JOIN order_items ON orders.id = order_items.order_id", id)
	err := row.Scan(&order.ID, &order.UserID, &order.Total, &order.Items, &order.CreatedAt, &order.UpdatedAt)
	return &order, err
}

package repository

import (
	"database/sql"

	"github.com/janczizikow/bookshelf/internal/models"
)

type booksRepository struct {
	db *sql.DB
}

type BooksRepository interface {
	Find(id int) (*models.Book, error)
}

// NewBooksRepository returns a new instance of a books repository.
func NewBooksRepository(db *sql.DB) *booksRepository {
	return &booksRepository{db: db}
}

func (r *booksRepository) Find(id int) (*models.Book, error) {
	book := models.Book{}
	row := r.db.QueryRow("SELECT * FROM books WHERE id = ?", id)
	err := row.Scan(&book.ID, &book.Title, &book.ISBN, &book.Price, &book.Publisher, &book.Pages, &book.PublishDate, &book.CreatedAt, &book.UpdatedAt, &book.AuthorID)
	return &book, err
}

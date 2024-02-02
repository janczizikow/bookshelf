package repository

import (
	"database/sql"

	"github.com/janczizikow/bookshelf/internal/models"
)

type authorsRepository struct {
	db *sql.DB
}

type AuthorsRepository interface {
	Find(id int) (*models.Author, error)
	FindByName(name string) (*models.Author, error)
	Create(author *models.Author) (*models.Author, error)
	Update(author *models.Author) (*models.Author, error)
	Delete(id int) error
}

// NewAuthorsRepository returns a new instance of a authors repository.
func NewAuthorsRepository(db *sql.DB) *authorsRepository {
	return &authorsRepository{db: db}
}

func (r *authorsRepository) Find(id int) (*models.Author, error) {
	author := models.Author{}
	row := r.db.QueryRow("SELECT * FROM authors WHERE id = ?", id)
	err := row.Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt)
	return &author, err
}

func (r *authorsRepository) FindByName(name string) (*models.Author, error) {
	author := models.Author{}
	row := r.db.QueryRow("SELECT * FROM authors WHERE name = ?", name)
	err := row.Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt)
	return &author, err
}

func (r *authorsRepository) Create(author *models.Author) (*models.Author, error) {
	newAuthor := models.Author{}
	query := `INSERT INTO authors (name)
						VALUES (?)
						RETURNING *`
	row := r.db.QueryRow(query, author.Name)
	err := row.Scan(&newAuthor.ID, &newAuthor.Name, &newAuthor.CreatedAt, &newAuthor.UpdatedAt)
	return &newAuthor, err
}

func (r *authorsRepository) Update(author *models.Author) (*models.Author, error) {
	updated := models.Author{}
	query := `UPDATE authors
						SET name = ?, updated_at = now()
						WHERE id = ?
						RETURNING *`
	row := r.db.QueryRow(query, author.Name, author.ID)
	err := row.Scan(&updated.ID, &updated.Name, &updated.CreatedAt, &updated.UpdatedAt)
	return &updated, err
}

func (r *authorsRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM authors WHERE id = ?", id)
	return err
}

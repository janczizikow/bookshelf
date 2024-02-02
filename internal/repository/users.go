package repository

import (
	"database/sql"

	"github.com/janczizikow/bookshelf/internal/models"
)

type usersRepository struct {
	db *sql.DB
}

type UserRepository interface {
	Find(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id int) error
}

// NewUsersRepository returns a new instance of a users repository.
func NewUsersRepository(db *sql.DB) *usersRepository {
	return &usersRepository{db: db}
}

func (r *usersRepository) Find(id int) (*models.User, error) {
	user := models.User{}
	row := r.db.QueryRow("SELECT * FROM users WHERE id = ? AND deleted_at IS NULL", id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	return &user, err
}

func (r *usersRepository) FindByEmail(email string) (*models.User, error) {
	user := models.User{}
	row := r.db.QueryRow("SELECT * FROM users WHERE email = ? AND deleted_at IS NULL", email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	return &user, err
}

func (r *usersRepository) Create(user *models.User) (*models.User, error) {
	newUser := models.User{}
	query := `INSERT INTO users (email, password_hash, name)
						VALUES (?, ?, ?)
						RETURNING *`
	row := r.db.QueryRow(query, user.Email, user.Password, user.Name)
	err := row.Scan(&newUser.ID, &newUser.Email, &newUser.Password, &newUser.Name, &newUser.CreatedAt, &newUser.UpdatedAt, &newUser.DeletedAt)
	return &newUser, err
}

func (r *usersRepository) Update(user *models.User) (*models.User, error) {
	updated := models.User{}
	// TODO:
	return &updated, nil
}

func (r *usersRepository) Delete(id int) error {
	_, err := r.db.Exec("UPDATE users SET deleted_at = now() WHERE id = ? AND deleted_at IS NULL", id)
	return err
}

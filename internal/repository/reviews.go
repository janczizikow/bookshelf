package repository

import (
	"database/sql"

	"github.com/janczizikow/bookshelf/internal/models"
)

type reviewsRepository struct {
	db *sql.DB
}

type ReviewsRepository interface {
	// TODO:
	// FindByBookID(bookID int) (error)
	Create(review *models.Review) (*models.Review, error)
	Update(review *models.Review) (*models.Review, error)
	Delete(id int) error
}

// NewReviewsRepository returns a new instance of a reviews repository.
func NewReviewsRepository(db *sql.DB) *reviewsRepository {
	return &reviewsRepository{db: db}
}

func (r *reviewsRepository) Create(review *models.Review) (*models.Review, error) {
	newReview := models.Review{}
	query := `INSERT INTO reviews (rating, description, book_id, user_id)
						RETURNING *`
	row := r.db.QueryRow(query, review.Rating, review.Description, review.BookID, review.UserID)
	err := row.Scan(&newReview.ID, &newReview.Rating, &newReview.Description, &newReview.BookID, &newReview.UserID, &newReview.CreatedAt, &newReview.UpdatedAt)
	return &newReview, err
}

func (r *reviewsRepository) Update(review *models.Review) (*models.Review, error) {
	updated := models.Review{}
	query := `UPDATE reviews
						SET rating = ?, description = ?, updated_at = now()
						WHERE id = ?
						RETURNING *`
	row := r.db.QueryRow(query, review.Rating, review.Description, review.ID)
	err := row.Scan(&updated.ID, &updated.Rating, &updated.Description, &updated.BookID, &updated.UserID, &updated.CreatedAt, &updated.UpdatedAt)
	return &updated, err
}

func (r *reviewsRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM reviews WHERE id = ?", id)
	return err
}

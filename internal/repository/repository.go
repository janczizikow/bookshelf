package repository

import "database/sql"

type Repository struct {
	db      *sql.DB
	Users   UserRepository
	Authors AuthorsRepository
	Books   BooksRepository
	Reviews ReviewsRepository
	Orders  OrdersRepository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db:      db,
		Users:   NewUsersRepository(db),
		Authors: NewAuthorsRepository(db),
		Books:   NewBooksRepository(db),
		Reviews: NewReviewsRepository(db),
		Orders:  NewOrdersRepository(db),
	}
}

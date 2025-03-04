package repository

import (
	"database/sql"

	commonrepository "github.com/tksasha/balance/internal/common/repository"
)

type Repository struct {
	*commonrepository.Repository

	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Repository: commonrepository.New(),
		db:         db,
	}
}

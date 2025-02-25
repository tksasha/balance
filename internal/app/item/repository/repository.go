package repository

import (
	"database/sql"

	"github.com/tksasha/balance/internal/common/repository"
)

type Repository struct {
	*repository.Repository

	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Repository: repository.New(),
		db:         db,
	}
}

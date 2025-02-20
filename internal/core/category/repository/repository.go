package repository

import (
	"database/sql"

	"github.com/tksasha/balance/internal/core/common/repositories"
)

type Repository struct {
	*repositories.BaseRepository

	db *sql.DB
}

func New(
	baseRepository *repositories.BaseRepository,
	db *sql.DB,
) *Repository {
	return &Repository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

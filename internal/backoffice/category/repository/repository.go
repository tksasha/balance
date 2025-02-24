package repository

import (
	"database/sql"

	"github.com/tksasha/balance/internal/common"
)

type Repository struct {
	*common.BaseRepository

	db *sql.DB
}

func New(baseRepository *common.BaseRepository, db *sql.DB) *Repository {
	return &Repository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetCategory(ctx context.Context, id int) (*models.Category, error) {
	query := `
		SELECT
			id,
			name
		FROM
			categories
		WHERE
			id=?
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var category models.Category

	if err := row.Scan(&category.ID, &category.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, internalerrors.ErrNotFound
		}

		slog.Error(err.Error())

		return nil, internalerrors.ErrUnknown
	}

	return &category, nil
}

package repositories

import (
	"context"
	"database/sql"
	"errors"

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

		panic(err)
	}

	return &category, nil
}

func (r *CategoryRepository) GetCategories(ctx context.Context, currency int) ([]*models.Category, error) {
	query := `
		SELECT
			categories.id,
			categories.name,
			categories.income
		FROM
			categories
		WHERE
			categories.visible=true
			AND
				categories.currency=?
		ORDER BY
			categories.name ASC
	`

	rows, err := r.db.QueryContext(ctx, query, currency)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()

	var categories []*models.Category

	for rows.Next() {
		var category models.Category

		if err := rows.Scan(&category.ID, &category.Name, &category.Income); err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

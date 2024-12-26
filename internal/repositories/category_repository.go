package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context) (models.Categories, error)
}

type categoryRepository struct {
	db *db.DB
}

func NewCategoryRepository(db *db.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetCategory(ctx context.Context, id int) (*models.Category, error) {
	query := `
		SELECT
			id,
			name
		FROM
			categories
		WHERE
			id=?
	`

	row := r.db.Connection.QueryRowContext(ctx, query, id)

	var category models.Category

	if err := row.Scan(&category.ID, &category.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewNotFoundError()
		}

		panic(err)
	}

	return &category, nil
}

func (r *categoryRepository) GetCategories(ctx context.Context) (models.Categories, error) {
	currency, ok := ctx.Value(models.CurrencyContextValue{}).(models.Currency)
	if !ok {
		currency, _ = models.GetDefaultCurrency()
	}

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

	rows, err := r.db.Connection.QueryContext(ctx, query, currency)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()

	var categories models.Categories

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

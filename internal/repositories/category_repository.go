package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
)

type CategoryRepository struct {
	db *db.DB
}

func NewCategoryRepository(db *db.DB) *CategoryRepository {
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

func (r *CategoryRepository) GetAll(ctx context.Context) (models.Categories, error) {
	currency, ok := ctx.Value(models.CurrencyContextValue{}).(models.Currency)
	if !ok {
		currency, _ = models.GetDefaultCurrency()
	}

	query := `
		SELECT
			id, name, income
		FROM
			categories
		WHERE
			visible=true AND categories.currency=?
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
		category := &models.Category{}

		if err := rows.Scan(&category.ID, &category.Name, &category.Income); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	currency, ok := ctx.Value(models.CurrencyContextValue{}).(models.Currency)
	if !ok {
		currency, _ = models.GetDefaultCurrency()
	}

	query := `INSERT INTO categories (name, income, visible, currency, supercategory) VALUES (?, ?, ?, ?, ?)`

	if _, err := r.db.Connection.ExecContext(
		ctx,
		query,
		category.Name,
		category.Income,
		category.Visible,
		currency,
		category.Supercategory,
	); err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) FindByName(ctx context.Context, name string) (*models.Category, error) {
	currency, ok := ctx.Value(models.CurrencyContextValue{}).(models.Currency)
	if !ok {
		currency, _ = models.GetDefaultCurrency()
	}

	query := `
		SELECT
			id,
			name,
			income,
			visible,
			currency,
			supercategory
		FROM
			categories
		WHERE
			name=? AND currency=?
		LIMIT 1
	`

	category := &models.Category{}

	err := r.db.Connection.
		QueryRowContext(
			ctx,
			query,
			name,
			currency,
		).
		Scan(
			&category.ID,
			&category.Name,
			&category.Income,
			&category.Visible,
			&category.Currency,
			&category.Supercategory,
		)
	if err != nil {
		return nil, err
	}

	return category, nil
}

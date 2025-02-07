package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/tksasha/balance/internal/apperrors"
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

func (r *CategoryRepository) FindByID(ctx context.Context, id int) (*models.Category, error) {
	currency := getCurrencyFromContext(ctx)

	query := `
		SELECT id, name, income, visible, currency, supercategory
		FROM categories
		WHERE id=? AND currency=?
	`

	row := r.db.QueryRowContext(ctx, query, id, currency)

	category := &models.Category{}

	if err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Income,
		&category.Visible,
		&category.Currency,
		&category.Supercategory,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrRecordNotFound
		}

		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) (models.Categories, error) {
	currency := getCurrencyFromContext(ctx)

	query := `
		SELECT id, name, income, currency
		FROM categories
		WHERE visible=true AND categories.currency=?
		ORDER BY categories.name ASC
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

	var categories models.Categories

	for rows.Next() {
		category := &models.Category{}

		if err := rows.Scan(&category.ID, &category.Name, &category.Income, &category.Currency); err != nil {
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
	currency := getCurrencyFromContext(ctx)

	query := `
		INSERT INTO categories (name, income, visible, currency, supercategory)
		VALUES (?, ?, ?, ?, ?)
	`

	if _, err := r.db.ExecContext(
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
	currency := getCurrencyFromContext(ctx)

	query := `
		SELECT id, name, currency
		FROM categories
		WHERE name=? AND currency=?
		LIMIT 1
	`

	category := &models.Category{}

	err := r.db.
		QueryRowContext(ctx, query, name, currency).
		Scan(&category.ID, &category.Name, &category.Currency)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrRecordNotFound
		}

		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	currency := getCurrencyFromContext(ctx)

	query := `
		UPDATE categories
		SET name=?, income=?, visible=?, supercategory=?
		WHERE id=? AND currency=?
	`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		category.Name,
		category.Income,
		category.Visible,
		category.Supercategory,
		category.ID,
		currency,
	); err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, category *models.Category) error {
	currency := getCurrencyFromContext(ctx)

	query := `UPDATE categories SET deleted_at=? WHERE id=? AND currency=?`

	if _, err := r.db.ExecContext(ctx, query, time.Now().UTC(), category.ID, currency); err != nil {
		return err
	}

	return nil
}

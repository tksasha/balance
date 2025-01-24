package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/tksasha/balance/internal/db"
	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *db.DB) *ItemRepository {
	return &ItemRepository{
		db: db.Connection,
	}
}

func (r *ItemRepository) GetItems(ctx context.Context) (models.Items, error) {
	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	query := `
		SELECT
			items.id,
			items.date,
			items.sum,
			COALESCE(items.category_name, ""),
			items.description
		FROM
			items
		WHERE
			items.currency=?
		AND
			items.deleted_at IS NULL
		ORDER BY
			items.date DESC
	`

	rows, err := r.db.QueryContext(ctx, query, currency)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			slog.Error("failed to close rows", "error", err)
		}
	}()

	var items models.Items

	for rows.Next() {
		item := &models.Item{}

		if err := rows.Scan(&item.ID, &item.Date, &item.Sum, &item.CategoryName, &item.Description); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) Create(ctx context.Context, item *models.Item) error {
	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	query := `
		INSERT INTO items (
			date,
			formula,
			sum,
			category_id,
			category_name,
			description,
			currency
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		item.Date.Format(time.DateOnly),
		item.Formula,
		item.Sum,
		item.CategoryID,
		item.CategoryName,
		item.Description,
		currency,
	); err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) FindByID(ctx context.Context, id int) (*models.Item, error) {
	query := `
		SELECT id, date, sum, category_id, category_name, description
		FROM items
		WHERE items.id=? AND items.deleted_at IS NULL
	`

	item := &models.Item{}

	if err := r.db.
		QueryRowContext(ctx, query, id).
		Scan(
			&item.ID,
			&item.Date,
			&item.Sum,
			&item.CategoryID,
			&item.CategoryName,
			&item.Description,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, internalerrors.ErrRecordNotFound
		}

		return nil, err
	}

	return item, nil
}

func (r *ItemRepository) Update(ctx context.Context, item *models.Item) error {
	query := `
		UPDATE items
		SET date=?, formula=?, sum=?, category_id=?, category_name=?, description=?
		WHERE id=?
	`

	if _, err := r.db.ExecContext(ctx, query,
		item.Date,
		item.Formula,
		item.Sum,
		item.CategoryID,
		item.CategoryName,
		item.Description,
		item.ID,
	); err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) DeleteItem(ctx context.Context, id int) error {
	return nil
}

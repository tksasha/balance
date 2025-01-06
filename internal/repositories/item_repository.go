package repositories

import (
	"context"
	"log/slog"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
)

type ItemRepository struct {
	db *db.DB
}

func NewItemRepository(db *db.DB) *ItemRepository {
	return &ItemRepository{db}
}

func (r *ItemRepository) GetItems(ctx context.Context) (models.Items, error) {
	currency, ok := ctx.Value(models.CurrencyContextValue{}).(models.Currency)
	if !ok {
		currency, _ = models.GetDefaultCurrency()
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

	rows, err := r.db.Connection.QueryContext(ctx, query, currency)
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

func (r *ItemRepository) CreateItem(ctx context.Context, item *models.Item) error {
	return nil
}

func (r *ItemRepository) GetItem(ctx context.Context, id int) (*models.Item, error) {
	return &models.Item{}, nil
}

func (r *ItemRepository) UpdateItem(ctx context.Context, item *models.Item) error {
	return nil
}

func (r *ItemRepository) DeleteItem(ctx context.Context, id int) error {
	return nil
}

package repositories

import (
	"context"
	"database/sql"

	"github.com/tksasha/balance/internal/models"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) GetItems(ctx context.Context) ([]models.Item, error) {
	query := `
		SELECT
			items.date,
			items.sum,
			categories.name,
			items.description
		FROM
			items
		INNER JOIN
			categories
		ON
			categories.id=items.category_id
		WHERE
			date BETWEEN "2024-09-01" AND "2024-09-30"
		AND
			items.currency=0
		ORDER BY
			items.date DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close() // TODO: need to think //nolint
	}()

	var items []models.Item

	for rows.Next() {
		var item models.Item

		if err := rows.Scan(&item.Date, &item.Sum, &item.Category, &item.Description); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		_ = err // TODO: log it
	}

	return items, nil
}

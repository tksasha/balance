package repositories

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
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
		if err := rows.Close(); err != nil {
			slog.Error(err.Error())
		}
	}()

	var items []models.Item

	for rows.Next() {
		var item models.Item

		if err := rows.Scan(&item.Date, &item.Sum, &item.CategoryName, &item.Description); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		slog.Error(err.Error())
	}

	return items, nil
}

func (r *ItemRepository) CreateItem(ctx context.Context, req *requests.CreateItemRequest) error {
	query := `
		INSERT INTO
			items (date, sum, category_id, description)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, req.Date, req.Sum, req.CategoryID, req.Description)
	if err != nil {
		return err
	}

	return nil
}

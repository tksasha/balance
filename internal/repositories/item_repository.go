package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	internalerrors "github.com/tksasha/balance/internal/errors"
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
			items.id,
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

		if err := rows.Scan(&item.ID, &item.Date, &item.Sum, &item.CategoryName, &item.Description); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		slog.Error(err.Error())
	}

	return items, nil
}

func (r *ItemRepository) CreateItem(
	ctx context.Context,
	item *models.Item,
) error {
	query := `
		INSERT INTO
			items (date, sum, category_id, description)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, item.Date, item.Sum, item.CategoryID, item.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) GetItem(ctx context.Context, id int) (*models.Item, error) {
	query := `
	SELECT
		id,
		date,
		sum,
		COALESCE(formula, ""),
		category_id,
		description
	FROM
		items
	WHERE
		id=?
	`

	var item models.Item

	row := r.db.QueryRowContext(ctx, query, id)

	if err := row.Scan(
		&item.ID,
		&item.Date,
		&item.Sum,
		&item.Formula,
		&item.CategoryID,
		&item.Description,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, internalerrors.NewNotFoundError(err)
		}

		return nil, internalerrors.NewUnknownError(err)
	}

	return &item, nil
}

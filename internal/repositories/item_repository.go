package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
)

type ItemRepository struct {
	db                 *sql.DB
	categoryRepository *CategoryRepository
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{
		db:                 db,
		categoryRepository: NewCategoryRepository(db),
	}
}

func (r *ItemRepository) GetItems(ctx context.Context, currency models.Currency) ([]*models.Item, error) {
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

	rows, err := r.db.QueryContext(ctx, query, currency.ID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			slog.Error(err.Error())
		}
	}()

	var items []*models.Item

	for rows.Next() {
		item := models.NewItem()

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

func (r *ItemRepository) CreateItem(ctx context.Context, item *models.Item) error {
	var err error

	item.CategoryName, err = r.getCategoryName(ctx, item.CategoryID)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO
			items (date, formula, sum, category_id, category_name, description)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = r.db.ExecContext(
		ctx,
		query,
		item.Date,
		item.Formula,
		item.Sum,
		item.CategoryID,
		item.CategoryName,
		item.Description,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) GetItem(ctx context.Context, id int) (*models.Item, error) {
	query := `
	SELECT
		items.id,
		items.date,
		items.sum,
		COALESCE(items.formula, ""),
		items.category_id,
		COALESCE(items.category_name, ""),
		items.description
	FROM
		items
	WHERE
		items.id=?
	`

	item := models.NewItem()

	row := r.db.QueryRowContext(ctx, query, id)

	if err := row.Scan(
		&item.ID,
		&item.Date,
		&item.Sum,
		&item.Formula,
		&item.CategoryID,
		&item.CategoryName,
		&item.Description,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, internalerrors.ErrNotFound
		}

		slog.Error(err.Error())

		return nil, internalerrors.ErrUnknown
	}

	return item, nil
}

func (r *ItemRepository) UpdateItem(ctx context.Context, item *models.Item) error {
	var err error

	item.CategoryName, err = r.getCategoryName(ctx, item.CategoryID)
	if err != nil {
		return err
	}

	query := `
		UPDATE
			items
		SET
			date=?,
			formula=?,
			sum=?,
			category_id=?,
			category_name=?,
			description=?
		WHERE
			id=?
	`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		item.Date,
		item.Formula,
		item.Sum,
		item.CategoryID,
		item.CategoryName,
		item.Description,
		item.ID,
	); err != nil {
		slog.Error(err.Error())

		return err
	}

	return nil
}

func (r *ItemRepository) DeleteItem(ctx context.Context, item *models.Item) error {
	query := `
		UPDATE
			items
		SET
			deleted_at=?
		WHERE
			id=?
	`

	if _, err := r.db.ExecContext(ctx, query, time.Now(), item.ID); err != nil {
		slog.Error(err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return internalerrors.ErrNotFound
		}

		return err
	}

	return nil
}

func (r *ItemRepository) getCategoryName(ctx context.Context, id int) (string, error) {
	category, err := r.categoryRepository.GetCategory(ctx, id)
	if err != nil {
		return "", err
	}

	return category.Name, nil
}

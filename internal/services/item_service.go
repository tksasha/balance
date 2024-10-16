package services

import (
	"context"
	"strconv"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type ItemService struct {
	itemsGetter repositories.ItemsGetter
	itemGetter  repositories.ItemGetter
	itemUpdater repositories.ItemUpdater
	itemCreator repositories.ItemCreator
	itemDeleter repositories.ItemDeleter
}

func NewItemService(
	itemsGetter repositories.ItemsGetter,
	itemGetter repositories.ItemGetter,
	itemUpdater repositories.ItemUpdater,
	itemCreator repositories.ItemCreator,
	itemDeleter repositories.ItemDeleter,
) *ItemService {
	return &ItemService{
		itemsGetter: itemsGetter,
		itemGetter:  itemGetter,
		itemUpdater: itemUpdater,
		itemCreator: itemCreator,
		itemDeleter: itemDeleter,
	}
}

func (s *ItemService) GetItems(ctx context.Context, currency models.Currency) (models.Items, error) {
	items, err := s.itemsGetter.GetItems(ctx, currency)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *ItemService) GetItem(ctx context.Context, input string) (*models.Item, error) {
	id, err := strconv.Atoi(input)
	if err != nil || id <= 0 {
		return nil, internalerrors.ErrNotFound
	}

	item, err := s.itemGetter.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemService) UpdateItem(ctx context.Context, item *models.Item) error {
	return s.itemUpdater.UpdateItem(ctx, item)
}

func (s *ItemService) CreateItem(ctx context.Context, item *models.Item) error {
	return s.itemCreator.CreateItem(ctx, item)
}

func (s *ItemService) DeleteItem(ctx context.Context, item *models.Item) error {
	return s.itemDeleter.DeleteItem(ctx, item)
}

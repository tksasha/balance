package services

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/models"
)

type ItemService struct {
	itemRepository ItemRepository
}

func NewItemService(itemRepository ItemRepository) *ItemService {
	return &ItemService{
		itemRepository: itemRepository,
	}
}

func (s *ItemService) GetItems(ctx context.Context) (models.Items, error) {
	items, err := s.itemRepository.GetItems(ctx)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *ItemService) GetItem(ctx context.Context, input string) (*models.Item, error) {
	id, err := strconv.Atoi(input)
	if err != nil || id <= 0 {
		return nil, NewNotFoundError()
	}

	item, err := s.itemRepository.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemService) UpdateItem(ctx context.Context, item *models.Item) error {
	return s.itemRepository.UpdateItem(ctx, item)
}

func (s *ItemService) CreateItem(ctx context.Context, item *models.Item) error {
	return s.itemRepository.CreateItem(ctx, item)
}

func (s *ItemService) DeleteItem(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil || id < 1 {
		return NewNotFoundError()
	}

	return s.itemRepository.DeleteItem(ctx, id)
}

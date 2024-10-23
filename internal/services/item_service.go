package services

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type ItemService interface {
	GetItems(ctx context.Context) (models.Items, error)
	GetItem(ctx context.Context, input string) (*models.Item, error)
	UpdateItem(ctx context.Context, item *models.Item) error
	CreateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, input string) error
}

type itemService struct {
	itemRepository repositories.ItemRepository
}

func NewItemService(itemRepository repositories.ItemRepository) ItemService {
	return &itemService{
		itemRepository: itemRepository,
	}
}

func (s *itemService) GetItems(ctx context.Context) (models.Items, error) {
	items, err := s.itemRepository.GetItems(ctx)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *itemService) GetItem(ctx context.Context, input string) (*models.Item, error) {
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

func (s *itemService) UpdateItem(ctx context.Context, item *models.Item) error {
	return s.itemRepository.UpdateItem(ctx, item)
}

func (s *itemService) CreateItem(ctx context.Context, item *models.Item) error {
	return s.itemRepository.CreateItem(ctx, item)
}

func (s *itemService) DeleteItem(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil || id < 1 {
		return NewNotFoundError()
	}

	return s.itemRepository.DeleteItem(ctx, id)
}

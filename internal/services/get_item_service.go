package services

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type GetItemService struct {
	itemGetter repositories.ItemGetter
}

func NewGetItemService(itemGetter repositories.ItemGetter) *GetItemService {
	return &GetItemService{
		itemGetter: itemGetter,
	}
}

func (s *GetItemService) GetItem(ctx context.Context, input string) (*models.Item, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, errors.NewInvalidError()
	}

	item, err := s.itemGetter.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

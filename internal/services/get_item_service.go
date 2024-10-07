package services

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/interfaces"
	"github.com/tksasha/balance/internal/models"
)

type GetItemService struct {
	itemGetter interfaces.ItemGetter
}

func NewGetItemService(itemGetter interfaces.ItemGetter) *GetItemService {
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

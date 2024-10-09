package services

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/decorators"
	internalerrors "github.com/tksasha/balance/internal/errors"
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

func (s *GetItemService) GetItem(ctx context.Context, input string) (*decorators.ItemDecorator, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, internalerrors.ErrNotFound
	}

	item, err := s.itemGetter.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return decorators.NewItemDecorator(item), nil
}

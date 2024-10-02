package services

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/interfaces"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
)

type GetItemService struct {
	itemGetter interfaces.ItemGetter
}

func NewGetItemService(app *app.App) *GetItemService {
	return &GetItemService{
		itemGetter: repositories.NewItemRepository(app.DB),
	}
}

func (s *GetItemService) GetItem(ctx context.Context, input string) (*models.Item, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, errors.NewNotFoundError(err)
	}

	item, err := s.itemGetter.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

package service

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/service"
	"github.com/tksasha/validation"
)

type Service struct {
	*service.Service

	itemRepository     item.Repository
	categoryRepository item.CategoryRepository
}

func New(
	itemRepository item.Repository,
	categoryRepository item.CategoryRepository,
) *Service {
	return &Service{
		Service:            service.New(),
		itemRepository:     itemRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *Service) setCategory(
	ctx context.Context,
	item *item.Item,
	categoryID string,
	validation *validation.Validation,
) error {
	item.CategoryID = validation.Integer("category", categoryID)

	if item.CategoryID == 0 {
		return nil
	}

	category, err := s.categoryRepository.FindByID(ctx, item.CategoryID)
	if err != nil {
		return err
	}

	item.CategoryName = category.Name

	return nil
}

func (s *Service) findByID(ctx context.Context, input string) (*item.Item, error) {
	id, err := strconv.Atoi(input)
	if err != nil || id <= 0 {
		return nil, common.ErrResourceNotFound
	}

	item, err := s.itemRepository.FindByID(ctx, id)
	if err != nil {
		return nil, s.MapError(err)
	}

	return item, nil
}

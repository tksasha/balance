package service

import (
	"context"

	"github.com/tksasha/balance/internal/category"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/pkg/validation"
)

type Service struct {
	itemRepository     item.Repository
	categoryRepository category.Repository
}

func New(itemRepository item.Repository, categoryRepository category.Repository) *Service {
	return &Service{
		itemRepository:     itemRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *Service) setCategory(
	ctx context.Context,
	item *item.Item,
	categoryID string,
	validate *validation.Validation,
) error {
	item.CategoryID = validate.Integer("category", categoryID)

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

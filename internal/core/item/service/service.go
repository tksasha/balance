package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/pkg/validation"
)

type Service struct {
	*common.BaseService

	itemRepository     item.Repository
	categoryRepository category.Repository
}

func New(
	baseService *common.BaseService,
	itemRepository item.Repository,
	categoryRepository category.Repository,
) *Service {
	return &Service{
		BaseService:        baseService,
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

	item.CategoryName = sql.NullString{String: category.Name, Valid: true}

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

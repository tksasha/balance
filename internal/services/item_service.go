package services

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validation"
)

type ItemService struct {
	itemRepository     ItemRepository
	categoryRepository CategoryRepository
}

func NewItemService(itemRepository ItemRepository, categoryRepository CategoryRepository) *ItemService {
	return &ItemService{
		itemRepository:     itemRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *ItemService) Create(ctx context.Context, request requests.CreateItemRequest) (*models.Item, error) {
	item := &models.Item{
		Formula:     request.Formula,
		Description: request.Description,
	}

	validate := validation.New()

	item.Date = validate.Date("date", request.Date)

	item.Sum = validate.Formula("sum", item.Formula)

	if err := s.setCategory(ctx, item, request.CategoryID, validate); err != nil {
		return nil, err
	}

	if validate.HasErrors() {
		return nil, validate.Errors
	}

	if err := s.itemRepository.Create(ctx, item); err != nil {
		return item, err
	}

	return item, nil
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
		return nil, apperrors.ErrResourceNotFound
	}

	item, err := s.itemRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemService) Update(ctx context.Context, request requests.UpdateItemRequest) error {
	item, err := s.findByID(ctx, request.ID)
	if err != nil {
		return apperrors.ErrResourceNotFound
	}

	validate := validation.New()

	item.Date = validate.Date("date", request.Date)

	item.Formula = request.Formula

	item.Sum = validate.Formula("sum", item.Formula)

	if err := s.setCategory(ctx, item, request.CategoryID, validate); err != nil {
		return err
	}

	item.Description = request.Description

	if validate.HasErrors() {
		return validate.Errors
	}

	return s.itemRepository.Update(ctx, item)
}

func (s *ItemService) Delete(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil || id < 1 {
		return apperrors.ErrResourceNotFound
	}

	return s.itemRepository.Delete(ctx, id)
}

func (s *ItemService) setCategory(
	ctx context.Context,
	item *models.Item,
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

func (s *ItemService) findByID(ctx context.Context, value string) (*models.Item, error) {
	id, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}

	return s.itemRepository.FindByID(ctx, id)
}

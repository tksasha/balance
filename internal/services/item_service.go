package services

import (
	"context"
	"strconv"
	"time"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validationerror"
	"github.com/tksasha/calculator"
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
	item := &models.Item{}

	validationErrors := validationerror.New()

	s.validateDate(request, item, validationErrors)
	s.validateFormula(request, item, validationErrors)
	s.validateCategory(ctx, request, item, validationErrors)

	item.Description = request.Description

	if validationErrors.Exists() {
		return item, validationErrors
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
		return nil, internalerrors.ErrResourceNotFound
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

func (s *ItemService) DeleteItem(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil || id < 1 {
		return internalerrors.ErrResourceNotFound
	}

	return s.itemRepository.DeleteItem(ctx, id)
}

func (s *ItemService) validateDate(
	request requests.CreateItemRequest,
	item *models.Item,
	validationErrors validationerror.ValidationError,
) {
	if request.Date == "" {
		validationErrors.Set("date", validationerror.IsRequired)

		return
	}

	date, err := time.Parse(time.DateOnly, request.Date)
	if err != nil {
		validationErrors.Set("date", validationerror.IsInvalid)
	}

	item.Date = date
}

func (s *ItemService) validateFormula(
	request requests.CreateItemRequest,
	item *models.Item,
	validationErrors validationerror.ValidationError,
) {
	if request.Formula == "" {
		validationErrors.Set("sum", validationerror.IsRequired)

		return
	}

	sum, err := calculator.Calculate(request.Formula)
	if err != nil {
		validationErrors.Set("sum", validationerror.IsInvalid)
	}

	item.Formula = request.Formula
	item.Sum = sum
}

func (s *ItemService) validateCategory(
	ctx context.Context,
	request requests.CreateItemRequest,
	item *models.Item,
	validationErrors validationerror.ValidationError,
) {
	if request.CategoryID == "" {
		validationErrors.Set("category_id", validationerror.IsRequired)

		return
	}

	categoryID, err := strconv.Atoi(request.CategoryID)
	if err != nil {
		validationErrors.Set("category_id", validationerror.IsInvalid)

		return
	}

	item.CategoryID = categoryID

	category, err := s.categoryRepository.FindByID(ctx, categoryID)
	if err != nil {
		validationErrors.Set("category_id", validationerror.InternalError)

		return
	}

	item.CategoryName = category.Name
}

package services

import (
	"context"
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validationerror"
	"github.com/tksasha/calculator"
)

type CreateItemService struct {
	itemRepository     ItemRepository
	categoryRepository CategoryRepository
}

func NewCreateItemService(itemRepository ItemRepository, categoryRepository CategoryRepository) *CreateItemService {
	return &CreateItemService{
		itemRepository:     itemRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *CreateItemService) Create(ctx context.Context, request requests.CreateItemRequest) (*models.Item, error) {
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

func (s *CreateItemService) validateDate(
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

func (s *CreateItemService) validateFormula(
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

func (s *CreateItemService) validateCategory(
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

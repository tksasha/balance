package services

import (
	"context"
	"errors"
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

	s.setDate(item, request.Date, validationErrors)
	s.setSum(item, request.Formula, validationErrors)

	if err := s.setCategory(ctx, item, request.CategoryID, validationErrors); err != nil {
		return item, err
	}

	s.setDescription(item, request.Description)

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

	item, err := s.itemRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemService) Update(ctx context.Context, request requests.UpdateItemRequest) error {
	id, err := strconv.Atoi(request.ID)
	if err != nil {
		return internalerrors.ErrResourceNotFound
	}

	item, err := s.itemRepository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, internalerrors.ErrRecordNotFound) {
			return internalerrors.ErrResourceNotFound
		}

		return err
	}

	validationErrors := validationerror.New()

	s.setDate(item, request.Date, validationErrors)
	s.setSum(item, request.Formula, validationErrors)

	if err := s.setCategory(ctx, item, request.CategoryID, validationErrors); err != nil {
		return err
	}

	s.setDescription(item, request.Description)

	if validationErrors.Exists() {
		return validationErrors
	}

	return s.itemRepository.Update(ctx, item)
}

func (s *ItemService) DeleteItem(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil || id < 1 {
		return internalerrors.ErrResourceNotFound
	}

	return s.itemRepository.DeleteItem(ctx, id)
}

func (s *ItemService) setDate(
	item *models.Item,
	value string,
	validationErrors validationerror.ValidationError,
) {
	if value == "" {
		validationErrors.Set("date", validationerror.IsRequired)

		return
	}

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		validationErrors.Set("date", validationerror.IsInvalid)
	}

	item.Date = date
}

func (s *ItemService) setSum(
	item *models.Item,
	value string,
	validationErrors validationerror.ValidationError,
) {
	if value == "" {
		validationErrors.Set("sum", validationerror.IsRequired)

		return
	}

	sum, err := calculator.Calculate(value)
	if err != nil {
		validationErrors.Set("sum", validationerror.IsInvalid)
	}

	item.Formula = value
	item.Sum = sum
}

func (s *ItemService) setCategory(
	ctx context.Context,
	item *models.Item,
	value string,
	validationErrors validationerror.ValidationError,
) error {
	if value == "" {
		validationErrors.Set("category", validationerror.IsRequired)

		return nil
	}

	categoryID, err := strconv.Atoi(value)
	if err != nil {
		validationErrors.Set("category", validationerror.IsInvalid)

		return nil //nolint:nilerr
	}

	item.CategoryID = categoryID

	category, err := s.categoryRepository.FindByID(ctx, categoryID)
	if err != nil {
		return err
	}

	item.CategoryName = category.Name

	return nil
}

func (s *ItemService) setDescription(
	item *models.Item,
	value string,
) {
	item.Description = value
}

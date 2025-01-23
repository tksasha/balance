package services

import (
	"context"
	"errors"
	"strconv"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validationerror"
)

type CategoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(categoryRepository CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Create(ctx context.Context, request requests.CreateCategoryRequest) error {
	category := &models.Category{
		Income:  request.Income == "true",
		Visible: request.Visible == "true",
	}

	validationErrors := validationerror.New()

	if err := s.validateName(ctx, request, category, validationErrors); err != nil {
		return err
	}

	s.validateSupercategory(request, category, validationErrors)

	if validationErrors.Exists() {
		return validationErrors
	}

	if err := s.categoryRepository.Create(ctx, category); err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) GetAll(ctx context.Context) (models.Categories, error) {
	return s.categoryRepository.GetAll(ctx)
}

func (s *CategoryService) FindByID(ctx context.Context, id int) (*models.Category, error) {
	category, err := s.categoryRepository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, internalerrors.ErrRecordNotFound) {
			return nil, internalerrors.ErrResourceNotFound
		}

		return nil, err
	}

	return category, nil
}

func (s *CategoryService) Update(ctx context.Context, categoryToUpdate *models.Category) error {
	validationError := validationerror.New()

	if categoryToUpdate.Name == "" {
		validationError.Set("name", validationerror.IsRequired)
	} else {
		categoryFound, err := s.categoryRepository.FindByName(ctx, categoryToUpdate.Name)
		if err != nil && !errors.Is(err, internalerrors.ErrRecordNotFound) {
			return err
		}

		if categoryFound != nil && categoryFound.ID != categoryToUpdate.ID {
			validationError.Set("name", validationerror.AlreadyExists)
		}
	}

	if validationError.Exists() {
		return validationError
	}

	return s.categoryRepository.Update(ctx, categoryToUpdate)
}

func (s *CategoryService) Delete(ctx context.Context, category *models.Category) error {
	return s.categoryRepository.Delete(ctx, category)
}

func (s *CategoryService) validateName(
	ctx context.Context,
	request requests.CreateCategoryRequest,
	category *models.Category,
	validationErrors validationerror.ValidationError,
) error {
	if request.Name == "" {
		validationErrors.Set("name", validationerror.IsRequired)

		return nil
	}

	category.Name = request.Name

	category, err := s.categoryRepository.FindByName(ctx, category.Name)
	if err != nil && !errors.Is(err, internalerrors.ErrRecordNotFound) {
		return err
	}

	if category != nil {
		validationErrors.Set("name", validationerror.AlreadyExists)
	}

	return nil
}

func (s *CategoryService) validateSupercategory(
	request requests.CreateCategoryRequest,
	category *models.Category,
	validationErrors validationerror.ValidationError,
) {
	if request.Supercategory == "" {
		return
	}

	supercategory, err := strconv.Atoi(request.Supercategory)
	if err != nil {
		validationErrors.Set("supercategory", validationerror.IsInvalid)
	}

	category.Supercategory = supercategory
}

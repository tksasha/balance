package services

import (
	"context"
	"errors"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
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

func (s *CategoryService) GetAll(ctx context.Context) (models.Categories, error) {
	return s.categoryRepository.GetAll(ctx)
}

func (s *CategoryService) Create(ctx context.Context, category *models.Category) error {
	validationError := validationerror.New()

	if category.Name == "" {
		validationError.Set("name", validationerror.IsRequired)
	} else {
		category, err := s.categoryRepository.FindByName(ctx, category.Name)
		if err != nil && !errors.Is(err, internalerrors.ErrRecordNotFound) {
			return err
		}

		if category != nil {
			validationError.Set("name", validationerror.AlreadyExists)
		}
	}

	if validationError.Exists() {
		return validationError
	}

	if err := s.categoryRepository.Create(ctx, category); err != nil {
		return err
	}

	return nil
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

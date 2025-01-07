package services

import (
	"context"

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

func (s *CategoryService) GetCategories(ctx context.Context) (models.Categories, error) {
	return s.categoryRepository.GetCategories(ctx)
}

func (s *CategoryService) Create(ctx context.Context, category *models.Category) error {
	validationError := validationerror.New()

	if category.Name == "" {
		validationError.Set("name", validationerror.IsRequired)
	}

	if _, err := s.categoryRepository.FindByName(ctx, category.Name); err == nil {
		validationError.Set("name", validationerror.AlreadyExists)
	}

	if validationError.Exists() {
		return validationError
	}

	if err := s.categoryRepository.Create(ctx, category); err != nil {
		return err
	}

	return nil
}

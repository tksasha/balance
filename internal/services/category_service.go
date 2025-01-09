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

func (s *CategoryService) GetAll(ctx context.Context) (models.Categories, error) {
	return s.categoryRepository.GetAll(ctx)
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

func (s *CategoryService) FindByID(ctx context.Context, id int) (*models.Category, error) {
	return s.categoryRepository.FindByID(ctx, id)
}

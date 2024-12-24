package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type CategoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) GetCategories(ctx context.Context) (models.Categories, error) {
	categories, err := s.categoryRepository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

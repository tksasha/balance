package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type CategoryService struct {
	categoriesGetter repositories.CategoriesGetter
}

func NewCategoryService(categoriesGetter repositories.CategoriesGetter) *CategoryService {
	return &CategoryService{
		categoriesGetter: categoriesGetter,
	}
}

func (s *CategoryService) GetCategories(ctx context.Context, currency models.Currency) (models.Categories, error) {
	categories, err := s.categoriesGetter.GetCategories(ctx, currency)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

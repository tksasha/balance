package services

import (
	"context"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/repositories"
)

type GetCategoriesService struct {
	categoriesGetter repositories.CategoriesGetter
}

func NewGetCategoriesService(categoriesGetter repositories.CategoriesGetter) CategoriesGetter {
	return &GetCategoriesService{
		categoriesGetter: categoriesGetter,
	}
}

func (s *GetCategoriesService) GetCategories(
	ctx context.Context,
	currency int,
) (*decorators.CategoriesDecorator, error) {
	categories, err := s.categoriesGetter.GetCategories(ctx, currency)
	if err != nil {
		return nil, err
	}

	return decorators.NewCategoriesDecorator(categories), nil
}

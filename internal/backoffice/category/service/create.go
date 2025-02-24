package service

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/validator"
)

func (s *Service) Create(ctx context.Context, request category.CreateRequest) (*category.Category, error) {
	category := &category.Category{
		Name: request.Name,
	}

	validate := validator.New()

	validate.Presence("name", category.Name)

	if err := s.nameAlreadyExists(ctx, category.Name, 0, validate); err != nil {
		return nil, err
	}

	category.Supercategory = validate.Integer("supercategory", request.Supercategory)

	category.Income = validate.Boolean("income", request.Income)

	category.Visible = validate.Boolean("visible", request.Visible)

	if validate.HasErrors() {
		return category, validate.Errors
	}

	if err := s.repository.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

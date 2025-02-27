package service

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/validation"
)

func (s *Service) Create(ctx context.Context, request category.CreateRequest) (*category.Category, error) {
	category := &category.Category{
		Name: request.Name,
	}

	validation := validation.New()

	validation.Presence("name", category.Name)

	if err := s.nameAlreadyExists(ctx, category.Name, 0, validation); err != nil {
		return nil, err
	}

	category.Supercategory = validation.Integer("supercategory", request.Supercategory)

	category.Income = validation.Boolean("income", request.Income)

	category.Visible = validation.Boolean("visible", request.Visible)

	if validation.Errors.Exists() {
		return category, validation.Errors
	}

	if err := s.repository.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

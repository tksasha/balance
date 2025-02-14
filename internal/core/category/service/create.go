package service

import (
	"context"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/pkg/validation"
)

func (s *Service) Create(ctx context.Context, request category.CreateRequest) error {
	category := &category.Category{
		Name: request.Name,
	}

	validate := validation.New()

	validate.Presence("name", category.Name)

	if err := s.nameAlreadyExists(ctx, category.Name, 0, validate); err != nil {
		return err
	}

	category.Supercategory = validate.Integer("supercategory", request.Supercategory)

	category.Income = validate.Boolean("income", request.Income)

	category.Visible = validate.Boolean("visible", request.Visible)

	if validate.HasErrors() {
		return validate.Errors
	}

	return s.repository.Create(ctx, category)
}

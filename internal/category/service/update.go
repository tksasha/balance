package service

import (
	"context"

	"github.com/tksasha/balance/internal/category"
	"github.com/tksasha/balance/pkg/validation"
)

func (s *Service) Update(ctx context.Context, request category.UpdateRequest) (*category.Category, error) {
	category, err := s.findByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	validate := validation.New()

	category.Name = validate.Presence("name", request.Name)

	if err := s.nameAlreadyExists(ctx, category.Name, category.ID, validate); err != nil {
		return nil, err
	}

	category.Income = validate.Boolean("income", request.Income)

	category.Visible = validate.Boolean("visible", request.Visible)

	category.Supercategory = validate.Integer("supercategory", request.Supercategory)

	if validate.HasErrors() {
		return category, validate.Errors
	}

	if err := s.repository.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

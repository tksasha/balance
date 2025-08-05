package service

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/validation"
	"github.com/tksasha/xstrings"
)

func (s *Service) Update(ctx context.Context, request category.UpdateRequest) (*category.Category, error) {
	category, err := s.findByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	validation := validation.New()

	category.Name = validation.Presence("name", request.Name)

	category.Slug = xstrings.Transliterate(category.Name)

	if err := s.nameAlreadyExists(ctx, category.Name, category.ID, validation); err != nil {
		return nil, err
	}

	category.Income = validation.Boolean("income", request.Income)

	category.Visible = validation.Boolean("visible", request.Visible)

	category.Supercategory = validation.Integer("supercategory", request.Supercategory)

	category.Number = validation.Integer("number", request.Number) // TODO: test me

	category.Currency = currency.GetByCode(request.Currency) // TODO: test me

	if validation.Errors.Exists() {
		return category, validation.Errors
	}

	if err := s.repository.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

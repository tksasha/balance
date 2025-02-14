package service

import (
	"context"
	"errors"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/services"
	"github.com/tksasha/balance/pkg/validation"
)

func (s *Service) Update(ctx context.Context, request cash.UpdateRequest) (*cash.Cash, error) {
	cash, err := s.findByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	validate := validation.New()

	cash.Formula, cash.Sum = validate.Formula("sum", request.Formula)

	cash.Name = validate.Presence("name", request.Name)

	if cash.Name != "" {
		exists, err := s.repository.NameExists(ctx, cash.Name, cash.ID)
		if err != nil {
			return nil, err
		}

		if exists {
			validate.Set("name", services.AlreadyExists)
		}
	}

	cash.Supercategory = validate.Integer("supercategory", request.Supercategory)

	cash.Favorite = validate.Boolean("favorite", request.Favorite)

	if validate.HasErrors() {
		return cash, validate.Errors
	}

	if err := s.repository.Update(ctx, cash); err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrResourceNotFound
		}

		return nil, err
	}

	return cash, nil
}

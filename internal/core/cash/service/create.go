package service

import (
	"context"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/pkg/validation"
)

func (s *Service) Create(ctx context.Context, request cash.CreateRequest) (*cash.Cash, error) {
	cash := &cash.Cash{}

	validate := validation.New()

	cash.Name = validate.Presence("name", request.Name)

	if cash.Name != "" {
		exists, err := s.repository.NameExists(ctx, cash.Name, 0)
		if err != nil {
			return nil, err
		}

		if exists {
			validate.Set("name", common.AlreadyExists)
		}
	}

	cash.Formula, cash.Sum = validate.Formula("formula", request.Formula)

	cash.Supercategory = validate.Integer("supercategory", request.Supercategory)

	cash.Favorite = validate.Boolean("favorite", request.Favorite)

	if validate.HasErrors() {
		return cash, validate.Errors
	}

	return cash, s.repository.Create(ctx, cash)
}

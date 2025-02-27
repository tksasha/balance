package service

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/validation"
)

func (s *Service) Create(ctx context.Context, request cash.CreateRequest) (*cash.Cash, error) {
	cash := &cash.Cash{}

	validation := validation.New()

	cash.Name = validation.Presence("name", request.Name)

	if cash.Name != "" {
		exists, err := s.repository.NameExists(ctx, cash.Name, 0)
		if err != nil {
			return nil, err
		}

		if exists {
			validation.Errors.Set("name", common.AlreadyExists)
		}
	}

	cash.Formula, cash.Sum = validation.Formula("formula", request.Formula)

	cash.Supercategory = validation.Integer("supercategory", request.Supercategory)

	cash.Favorite = validation.Boolean("favorite", request.Favorite)

	if validation.Errors.Exists() {
		return cash, validation.Errors
	}

	return cash, s.repository.Create(ctx, cash)
}

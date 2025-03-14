package service

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/validation"
)

func (s *Service) Create(ctx context.Context, request cash.CreateRequest) (*cash.Cash, error) {
	cash := &cash.Cash{}

	validation := validation.New()

	cash.Name = validation.Presence("name", request.Name)

	cash.Currency = currency.GetByCode(request.Currency)

	if cash.Name != "" {
		exists, err := s.repository.NameExists(ctx, cash)
		if err != nil {
			return nil, err
		}

		if exists {
			validation.Errors.Set("name", common.AlreadyExists)
		}
	}

	cash.Formula, cash.Sum = validation.Formula("sum", request.Formula)

	cash.Supercategory = validation.Integer("supercategory", request.Supercategory)

	if validation.Errors.Exists() {
		return cash, validation.Errors
	}

	return cash, s.repository.Create(ctx, cash)
}

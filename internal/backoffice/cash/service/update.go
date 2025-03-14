package service

import (
	"context"
	"errors"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/validation"
)

func (s *Service) Update(ctx context.Context, request cash.UpdateRequest) (*cash.Cash, error) {
	cash, err := s.findByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

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

	if err := s.repository.Update(ctx, cash); err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrResourceNotFound
		}

		return nil, err
	}

	return cash, nil
}

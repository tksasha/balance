package service

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/tksasha/balance/internal/app/balance"
)

func (s *Service) Balance(ctx context.Context) (*balance.Balance, error) {
	atEnd, err := s.atEnd(ctx)
	if err != nil {
		return nil, err
	}

	cashes, err := s.repository.Cashes(ctx)
	if err != nil {
		return nil, err
	}

	return &balance.Balance{
		AtEnd:   atEnd,
		Balance: cashes.Sub(atEnd),
	}, nil
}

func (s *Service) atEnd(ctx context.Context) (decimal.Decimal, error) {
	income, err := s.repository.Income(ctx)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	expense, err := s.repository.Expense(ctx)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	return income.Sub(expense), nil
}

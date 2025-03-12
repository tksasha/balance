package service

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/tksasha/balance/internal/app/balance"
)

const cents = 2

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
		Balance: sub(cashes, atEnd),
	}, nil
}

func (s *Service) atEnd(ctx context.Context) (float64, error) {
	income, err := s.repository.Income(ctx)
	if err != nil {
		return 0, err
	}

	expense, err := s.repository.Expense(ctx)
	if err != nil {
		return 0, err
	}

	return sub(income, expense), nil
}

func sub(l, r float64) float64 {
	sub, _ := decimal.NewFromFloat(l).Round(cents).Sub(
		decimal.NewFromFloat(r).Round(cents),
	).Round(cents).Float64()

	return sub
}

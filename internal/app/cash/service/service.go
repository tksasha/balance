package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/common"
)

type Service struct {
	repository cash.Repository
}

func New(repository cash.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) findByID(ctx context.Context, input string) (*cash.Cash, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, common.ErrResourceNotFound
	}

	cash, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrResourceNotFound
		}

		return nil, err
	}

	return cash, nil
}

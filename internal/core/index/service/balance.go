package service

import (
	"context"
)

func (s *Service) Balance(ctx context.Context) (float64, error) {
	residual, err := s.residual(ctx)
	if err != nil {
		return 0.0, err
	}

	cashes, err := s.repository.Cashes(ctx)
	if err != nil {
		return 0.0, err
	}

	return cashes - residual, nil
}

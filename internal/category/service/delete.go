package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
)

func (s *Service) Delete(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil {
		return apperrors.ErrResourceNotFound
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, apperrors.ErrRecordNotFound) {
			return apperrors.ErrResourceNotFound
		}

		return err
	}

	return nil
}

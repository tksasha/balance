package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/tksasha/balance/internal/common"
)

func (s *Service) Delete(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil {
		return common.ErrResourceNotFound
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return common.ErrResourceNotFound
		}

		return err
	}

	return nil
}

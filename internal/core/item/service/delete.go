package service

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/common"
)

func (s *Service) Delete(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil || id < 1 {
		return common.ErrResourceNotFound
	}

	return s.MapError(
		s.itemRepository.Delete(ctx, id),
	)
}

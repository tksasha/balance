package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/cash"
)

func (s *Service) GroupedList(ctx context.Context) (cash.GroupedCashes, error) {
	cashes, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	groupedCashes := cash.GroupedCashes{}

	for _, cash := range cashes {
		groupedCashes[cash.Supercategory] = append(groupedCashes[cash.Supercategory], cash)
	}

	return groupedCashes, nil
}

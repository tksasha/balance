package service

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/app/categoryreport"
	"github.com/tksasha/month"
)

func (s *Service) Report(ctx context.Context, request categoryreport.Request) (categoryreport.MappedEntities, error) {
	month := month.New(request.Year, request.Month)

	filters := categoryreport.Filters{
		From: month.Begin.Format(time.DateOnly),
		To:   month.End.Format(time.DateOnly),
	}

	entities, err := s.repository.Group(ctx, filters)
	if err != nil {
		return nil, err
	}

	mapped := categoryreport.MappedEntities{}

	for _, entity := range entities {
		mapped[entity.Supercategory] = append(mapped[entity.Supercategory], entity)
	}

	return mapped, nil
}

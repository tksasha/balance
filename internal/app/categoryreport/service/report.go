package service

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/app/categoryreport"
	"github.com/tksasha/month"
)

func (s *Service) Report(ctx context.Context, request categoryreport.Request) (categoryreport.Entities, error) {
	month := month.New(request.Year, request.Month)

	filters := categoryreport.Filters{
		From: month.Begin.Format(time.DateOnly),
		To:   month.End.Format(time.DateOnly),
	}

	return s.repository.Group(ctx, filters)
}

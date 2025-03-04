package service

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/month"
)

func (s *Service) List(ctx context.Context, request item.ListRequest) (item.Items, error) {
	month := month.New(request.Year, request.Month)

	filters := item.Filters{
		From: month.Begin.Format(time.DateOnly),
		To:   month.End.Format(time.DateOnly),
	}

	return s.itemRepository.FindAll(ctx, filters)
}

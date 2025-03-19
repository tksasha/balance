package service

import (
	"context"
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/month"
)

func (s *Service) List(ctx context.Context, request item.ListRequest) (item.Items, error) {
	month := month.New(request.Year, request.Month)

	category, _ := strconv.Atoi(request.Category)

	filters := item.Filters{
		From:     month.Begin.Format(time.DateOnly),
		To:       month.End.Format(time.DateOnly),
		Category: category,
	}

	return s.itemRepository.FindAll(ctx, filters)
}

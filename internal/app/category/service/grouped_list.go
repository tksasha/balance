package service

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/month"
)

func (s *Service) GroupedList(ctx context.Context, request category.Request) (category.GroupedCategories, error) {
	month := month.New(request.Year, request.Month)

	filters := category.Filters{
		From: month.Begin.Format(time.DateOnly),
		To:   month.End.Format(time.DateOnly),
	}

	categories, err := s.repository.FindAllByFilters(ctx, filters)
	if err != nil {
		return nil, err
	}

	grouped := category.GroupedCategories{}

	for _, category := range categories {
		grouped[category.Supercategory] = append(grouped[category.Supercategory], category)
	}

	return grouped, nil
}

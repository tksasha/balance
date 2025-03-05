package category

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) (Categories, error)
	FindAllByFilters(ctx context.Context, filters Filters) (Categories, error)
}

type Service interface {
	GroupedList(ctx context.Context, request Request) (GroupedCategories, error)
}

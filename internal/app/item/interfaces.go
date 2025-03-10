package item

import (
	"context"

	"github.com/tksasha/balance/internal/app/category"
)

type Repository interface {
	FindAll(ctx context.Context, filters Filters) (Items, error)
	Create(ctx context.Context, item *Item) error
	FindByID(ctx context.Context, id int) (*Item, error)
	Update(ctx context.Context, item *Item) error
	Delete(ctx context.Context, id int) error
}

type Service interface {
	Create(ctx context.Context, request CreateRequest) (*Item, error)
	List(ctx context.Context, request ListRequest) (Items, error)
	Edit(ctx context.Context, input string) (*Item, error)
	Update(ctx context.Context, request UpdateRequest) (*Item, error)
	Delete(ctx context.Context, input string) (*Item, error)
}

type CategoryRepository interface {
	FindByID(ctx context.Context, id int) (*category.Category, error)
}

type CategoryService interface {
	List(ctx context.Context) (category.Categories, error)
}

package item

import (
	"context"

	"github.com/tksasha/month"
)

type Repository interface {
	FindAllByMonth(ctx context.Context, month month.Month) (Items, error)
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
	Delete(ctx context.Context, input string) error
}

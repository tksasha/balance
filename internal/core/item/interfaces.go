package item

import (
	"context"
)

type Repository interface {
	List(ctx context.Context) (Items, error)
	Create(ctx context.Context, item *Item) error
	FindByID(ctx context.Context, id int) (*Item, error)
	Update(ctx context.Context, item *Item) error
	Delete(ctx context.Context, id int) error
}

type Service interface {
	Create(ctx context.Context, request CreateRequest) (*Item, error)
	List(ctx context.Context) (Items, error)
	FindByID(ctx context.Context, input string) (*Item, error)
	Update(ctx context.Context, request UpdateRequest) (*Item, error)
	Delete(ctx context.Context, input string) error
}

package category

import (
	"context"
)

type Repository interface {
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) (Categories, error)
	FindByName(ctx context.Context, name string) (*Category, error)
	Create(ctx context.Context, category *Category) error
	FindByID(ctx context.Context, id int) (*Category, error)
	Update(ctx context.Context, category *Category) error
}

type Service interface {
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) (Categories, error)
	Create(ctx context.Context, request CreateRequest) error
	FindByID(ctx context.Context, id string) (*Category, error)
	Update(ctx context.Context, request UpdateRequest) (*Category, error)
}

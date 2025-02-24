package category

import (
	"context"
)

type Repository interface {
	List(ctx context.Context) (Categories, error)
	FindByName(ctx context.Context, name string) (*Category, error)
	FindByID(ctx context.Context, id int) (*Category, error)
	Update(ctx context.Context, category *Category) error
}

type Service interface {
	List(ctx context.Context) (Categories, error)
	Edit(ctx context.Context, id string) (*Category, error)
	Update(ctx context.Context, request UpdateRequest) (*Category, error)
}

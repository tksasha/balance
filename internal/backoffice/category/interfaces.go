package category

import "context"

type Repository interface {
	FindAll(ctx context.Context) (Categories, error)
	FindByName(ctx context.Context, name string) (*Category, error)
	Create(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id int) error
}

type Service interface {
	List(ctx context.Context) (Categories, error)
	Create(ctx context.Context, request CreateRequest) (*Category, error)
	Delete(ctx context.Context, id string) error
}

package category

import "context"

type Repository interface {
	Create(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) (Categories, error)
	FindByID(ctx context.Context, id int) (*Category, error)
	FindByName(ctx context.Context, name string) (*Category, error)
	Update(ctx context.Context, category *Category) error
}

type Service interface {
	Create(ctx context.Context, request CreateRequest) (*Category, error)
	Delete(ctx context.Context, id string) error
	Edit(ctx context.Context, id string) (*Category, error)
	List(ctx context.Context) (Categories, error)
	Update(ctx context.Context, request UpdateRequest) (*Category, error)
}

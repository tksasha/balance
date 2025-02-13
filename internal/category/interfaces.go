package category

import (
	"context"
)

type Repository interface {
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) (Categories, error)
}

type Service interface {
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) (Categories, error)
}

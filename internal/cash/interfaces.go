package cash

import (
	"context"
)

type Repository interface {
	List(ctx context.Context) (Cashes, error)
	NameExists(ctx context.Context, name string, id int) (bool, error)
	Create(ctx context.Context, cash *Cash) error
}

type Service interface {
	List(ctx context.Context) (Cashes, error)
	Create(ctx context.Context, request CreateRequest) error
}

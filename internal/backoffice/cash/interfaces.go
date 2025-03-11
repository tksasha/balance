package cash

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) (Cashes, error)
	NameExists(ctx context.Context, name string, id int) (bool, error)
	Create(ctx context.Context, cash *Cash) error
	FindByID(ctx context.Context, id int) (*Cash, error)
	Update(ctx context.Context, cash *Cash) error
	Delete(ctx context.Context, id int) error
}

type Service interface {
	List(ctx context.Context) (Cashes, error)
	Create(ctx context.Context, request CreateRequest) (*Cash, error)
	Edit(ctx context.Context, id string) (*Cash, error)
	Update(ctx context.Context, request UpdateRequest) (*Cash, error)
	Delete(ctx context.Context, id string) error
}

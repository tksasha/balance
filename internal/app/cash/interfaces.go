package cash

import (
	"context"
)

type Repository interface {
	NameExists(ctx context.Context, name string, id int) (bool, error)
	FindByID(ctx context.Context, id int) (*Cash, error)
	Update(ctx context.Context, cash *Cash) error
	FindAll(ctx context.Context) (Cashes, error)
}

type Service interface {
	Edit(ctx context.Context, id string) (*Cash, error)
	Update(ctx context.Context, request UpdateRequest) (*Cash, error)
	List(ctx context.Context) (Cashes, error)
}

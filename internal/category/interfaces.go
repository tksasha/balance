package category

import (
	"context"
)

type Repository interface {
	Delete(ctx context.Context, id int) error
}

type Service interface {
	Delete(ctx context.Context, id string) error
}

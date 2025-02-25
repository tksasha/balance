package category

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) (Categories, error)
}

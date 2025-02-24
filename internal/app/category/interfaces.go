package category

import (
	"context"
)

type Repository interface {
	List(ctx context.Context) (Categories, error)
}

package index

import (
	"context"

	"github.com/tksasha/balance/internal/app/category"
)

type CategoryService interface {
	List(ctx context.Context) (category.Categories, error)
}

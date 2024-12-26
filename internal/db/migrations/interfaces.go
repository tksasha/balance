package migrations

import (
	"context"
)

type Migration interface {
	Up(ctx context.Context)
}

package cash

import (
	"context"
)

type Repository interface { //nolint:iface // tmp
	List(ctx context.Context) (Cashes, error)
}

type Service interface { //nolint:iface // tmp
	List(ctx context.Context) (Cashes, error)
}

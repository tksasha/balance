package categoryreport

import "context"

type Repository interface {
	Group(ctx context.Context, filters Filters) (Entities, error)
}

type Service interface {
	Report(ctx context.Context, request Request) (MappedEntities, error)
}

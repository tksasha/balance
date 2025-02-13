package service

import (
	"context"

	"github.com/tksasha/balance/internal/cash"
	"github.com/tksasha/balance/internal/common/services"
	"github.com/tksasha/balance/pkg/validation"
)

type Service struct {
	repository cash.Repository
}

func New(repository cash.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) List(ctx context.Context) (cash.Cashes, error) {
	return s.repository.List(ctx)
}

func (s *Service) Create(ctx context.Context, request cash.CreateRequest) error {
	cash := &cash.Cash{}

	validate := validation.New()

	cash.Name = validate.Presence("name", request.Name)

	if cash.Name != "" {
		exists, err := s.repository.NameExists(ctx, cash.Name, 0)
		if err != nil {
			return err
		}

		if exists {
			validate.Set("name", services.AlreadyExists)
		}
	}

	cash.Formula, cash.Sum = validate.Formula("formula", request.Formula)

	cash.Supercategory = validate.Integer("supercategory", request.Supercategory)

	cash.Favorite = validate.Boolean("favorite", request.Favorite)

	if validate.HasErrors() {
		return validate.Errors
	}

	return s.repository.Create(ctx, cash)
}

package service

import (
	"github.com/tksasha/balance/internal/core/indexpage"
)

type Service struct {
	repository indexpage.Repository
}

func New(repository indexpage.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

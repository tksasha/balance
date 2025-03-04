package service

import (
	"github.com/tksasha/balance/internal/app/categoryreport"
)

type Service struct {
	repository categoryreport.Repository
}

func New(repository categoryreport.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

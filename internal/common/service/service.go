package service

import (
	"errors"

	"github.com/tksasha/balance/internal/common"
)

const AlreadyExists = "already exists"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) MapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, common.ErrRecordNotFound) {
		return common.ErrResourceNotFound
	}

	return err
}

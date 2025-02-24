package common

import (
	"errors"
)

const AlreadyExists = "already exists"

type BaseService struct{}

func NewBaseService() *BaseService {
	return &BaseService{}
}

func (s *BaseService) MapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, ErrRecordNotFound) {
		return ErrResourceNotFound
	}

	return err
}

package services

import (
	"errors"

	"github.com/tksasha/balance/internal/core/common"
)

const AlreadyExists = "already exists"

func MapError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, common.ErrRecordNotFound) {
		return common.ErrResourceNotFound
	}

	return err
}

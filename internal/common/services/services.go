package services

import (
	"errors"

	"github.com/tksasha/balance/internal/core/common"
)

const AlreadyExists = "already exists"

func E(err error) error {
	if errors.Is(err, common.ErrRecordNotFound) {
		return common.ErrResourceNotFound
	}

	return err
}

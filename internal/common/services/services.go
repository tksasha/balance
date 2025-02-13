package services

import (
	"errors"

	"github.com/tksasha/balance/internal/apperrors"
)

const AlreadyExists = "already exists"

func E(err error) error {
	if errors.Is(err, apperrors.ErrRecordNotFound) {
		return apperrors.ErrResourceNotFound
	}

	return err
}

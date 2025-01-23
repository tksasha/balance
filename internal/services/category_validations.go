package services

import (
	"context"
	"errors"
	"strconv"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validationerror"
)

func validateCategoryName(
	ctx context.Context,
	categoryRepository CategoryRepository,
	request requests.CreateCategoryRequest,
	category *models.Category,
	validationErrors validationerror.ValidationError,
) error {
	if request.Name == "" {
		validationErrors.Set("name", validationerror.IsRequired)

		return nil
	}

	category.Name = request.Name

	category, err := categoryRepository.FindByName(ctx, category.Name)
	if err != nil && !errors.Is(err, internalerrors.ErrRecordNotFound) {
		return err
	}

	if category != nil {
		validationErrors.Set("name", validationerror.AlreadyExists)
	}

	return nil
}

func validateCategorySupercategory(
	request requests.CreateCategoryRequest,
	category *models.Category,
	validationErrors validationerror.ValidationError,
) {
	if request.Supercategory == "" {
		return
	}

	supercategory, err := strconv.Atoi(request.Supercategory)
	if err != nil {
		validationErrors.Set("supercategory", validationerror.IsInvalid)
	}

	category.Supercategory = supercategory
}

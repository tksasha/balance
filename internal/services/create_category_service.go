package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validationerror"
)

type CreateCategoryService struct {
	categoryRepository CategoryRepository
}

func NewCreateCategoryService(categoryRepository CategoryRepository) *CreateCategoryService {
	return &CreateCategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CreateCategoryService) Create(ctx context.Context, request requests.CreateCategoryRequest) error {
	category := &models.Category{
		Income:  request.Income == "true",
		Visible: request.Visible == "true",
	}

	validationErrors := validationerror.New()

	if err := validateCategoryName(ctx, s.categoryRepository, request, category, validationErrors); err != nil {
		return err
	}

	validateCategorySupercategory(request, category, validationErrors)

	if validationErrors.Exists() {
		return validationErrors
	}

	if err := s.categoryRepository.Create(ctx, category); err != nil {
		return err
	}

	return nil
}

package services

import (
	"context"
	"errors"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validation"
)

type CategoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(categoryRepository CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Create(ctx context.Context, request requests.CreateCategoryRequest) error {
	category := &models.Category{
		Name: request.Name,
	}

	validate := validation.New()

	validate.Presence("name", category.Name)

	if err := s.nameAlreadyExists(ctx, category.Name, 0, validate); err != nil {
		return err
	}

	category.Supercategory = validate.Integer("supercategory", request.Supercategory)

	category.Income = validate.Boolean("income", request.Income)

	category.Visible = validate.Boolean("visible", request.Visible)

	if validate.HasErrors() {
		return validate.Errors
	}

	return s.categoryRepository.Create(ctx, category)
}

func (s *CategoryService) GetAll(ctx context.Context) (models.Categories, error) {
	return s.categoryRepository.GetAll(ctx)
}

func (s *CategoryService) FindByID(ctx context.Context, id int) (*models.Category, error) {
	category, err := s.categoryRepository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, internalerrors.ErrRecordNotFound) {
			return nil, internalerrors.ErrResourceNotFound
		}

		return nil, err
	}

	return category, nil
}

func (s *CategoryService) Update(ctx context.Context, category *models.Category) error {
	validate := validation.New()

	validate.Presence("name", category.Name)

	if err := s.nameAlreadyExists(ctx, category.Name, category.ID, validate); err != nil {
		return err
	}

	if validate.HasErrors() {
		return validate.Errors
	}

	return s.categoryRepository.Update(ctx, category)
}

func (s *CategoryService) Delete(ctx context.Context, category *models.Category) error {
	return s.categoryRepository.Delete(ctx, category)
}

func (s *CategoryService) nameAlreadyExists(
	ctx context.Context,
	name string,
	categoryID int,
	validation *validation.Validation,
) error {
	if name == "" {
		return nil
	}

	category, err := s.categoryRepository.FindByName(ctx, name)

	if errors.Is(err, internalerrors.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return err
	}

	if category.ID != categoryID {
		validation.Set("name", AlreadyExists)
	}

	return nil
}

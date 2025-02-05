package handlers

import (
	"context"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
)

type ItemService interface {
	Create(ctx context.Context, request requests.CreateItemRequest) (*models.Item, error)
	GetItems(ctx context.Context) (models.Items, error)
	GetItem(ctx context.Context, input string) (*models.Item, error)
	Update(ctx context.Context, request requests.UpdateItemRequest) error
	Delete(ctx context.Context, input string) error
}

type CategoryService interface {
	Create(ctx context.Context, request requests.CreateCategoryRequest) error
	GetAll(ctx context.Context) (models.Categories, error)
	FindByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, category *models.Category) error
}

type CashService interface {
	Create(ctx context.Context, request requests.CreateCashRequest) error
}

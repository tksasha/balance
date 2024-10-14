package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/server/app"
)

type Handlers struct {
	Index      http.Handler
	GetItems   http.Handler
	CreateItem http.Handler
	EditItem   http.Handler
	UpdateItem http.Handler
	DeleteItem http.Handler
}

func New(app *app.App) *Handlers {
	return &Handlers{
		Index:      NewIndexHandler(app),
		GetItems:   NewGetItemsHandler(app),
		CreateItem: NewCreateItemHandler(app),
		EditItem:   NewEditItemHandler(app),
		UpdateItem: NewUpdateItemHandler(app),
		DeleteItem: NewDeleteItemHandler(app),
	}
}

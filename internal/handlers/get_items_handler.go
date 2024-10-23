package handlers

import "net/http"

type getItemsHandler struct{}

func NewGetItemsHandler() Route {
	return &getItemsHandler{}
}

func (h *getItemsHandler) Pattern() string {
	return "GET /items"
}

func (h *getItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

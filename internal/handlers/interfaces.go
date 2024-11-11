package handlers

import "net/http"

type Handler interface {
	http.Handler

	Pattern() string
}

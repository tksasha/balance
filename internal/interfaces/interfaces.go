package interfaces

import "net/http"

type Route interface {
	http.Handler

	Pattern() string
}

type Middleware interface {
	Wrap(next http.Handler) http.Handler
}

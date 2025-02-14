package middlewares

import "net/http"

type Middleware interface {
	Wrap(next http.Handler) http.Handler
}

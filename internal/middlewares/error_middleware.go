package middlewares

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/responses"
)

type errorMiddleware struct{}

func newErrorMiddleware() *errorMiddleware {
	return &errorMiddleware{}
}

func (m *errorMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responses.Response{ResponseWriter: w}

		next.ServeHTTP(rw, r)

		if err := rw.Error; err != nil {
			switch {
			case errors.Is(err, apperrors.ErrParsingForm):
				http.Error(w, "Bad Request", http.StatusBadRequest)
			case errors.Is(err, apperrors.ErrResourceNotFound):
				http.Error(w, "Resource Not Found", http.StatusNotFound)
			default:
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	})
}

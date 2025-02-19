package middlewares

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/response"
)

type errorMiddleware struct{}

func (m *errorMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, ok := w.(*response.Response)
		if !ok {
			slog.Error("invalid response writer", "response", response)
		}

		next.ServeHTTP(response, r)

		err := response.Error
		if err != nil {
			switch {
			case errors.Is(err, common.ErrParsingForm):
				response.Code = http.StatusBadRequest

				http.Error(response, "Bad Request", response.Code)
			case errors.Is(err, common.ErrResourceNotFound):
				response.Code = http.StatusNotFound

				http.Error(response, "Resource Not Found", response.Code)
			default:
				response.Code = http.StatusInternalServerError

				http.Error(response, "Internal Server Error", response.Code)
			}
		}
	})
}

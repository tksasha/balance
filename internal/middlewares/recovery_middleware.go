package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
)

type RecoveryMiddleware struct{}

func NewRecoveryMiddleware() *RecoveryMiddleware {
	return &RecoveryMiddleware{}
}

func (m *RecoveryMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("panic", "error", fmt.Sprintf("%v", err))

				fmt.Printf("%s\n", debug.Stack()) //nolint:forbidigo

				http.Error(w, "something went wrong", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

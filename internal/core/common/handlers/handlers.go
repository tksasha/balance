package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/response"
)

func SetError(w http.ResponseWriter, err error) {
	if wrapper, ok := w.(*response.Wrapper); ok {
		wrapper.Error = err
	}
}

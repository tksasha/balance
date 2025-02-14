package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/response"
)

func SetError(w http.ResponseWriter, err error) {
	if response, ok := w.(*response.Response); ok {
		response.Error = err
	}
}

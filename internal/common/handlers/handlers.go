package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/responses"
)

func E(w http.ResponseWriter, err error) {
	if response, ok := w.(*responses.Response); ok {
		response.Error = err
	}
}

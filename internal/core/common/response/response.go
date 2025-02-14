package response

import "net/http"

type Response struct {
	http.ResponseWriter

	Error error
}

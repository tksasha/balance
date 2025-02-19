package response

import "net/http"

type Response struct {
	http.ResponseWriter

	Code  int
	Error error
}

package responses

import "net/http"

type Response struct {
	http.ResponseWriter

	Error error
}

package response

import "net/http"

type Wrapper struct {
	http.ResponseWriter

	Code  int
	Error error
}

func NewWrapper(w http.ResponseWriter) *Wrapper {
	return &Wrapper{
		ResponseWriter: w,
		Code:           http.StatusOK,
	}
}

func (w *Wrapper) WriteHeader(code int) {
	w.Code = code

	w.ResponseWriter.WriteHeader(code)
}

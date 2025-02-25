package middleware

import "net/http"

type ResponseWriterWrapper struct {
	http.ResponseWriter

	Code  int
	Error error
}

func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{
		ResponseWriter: w,
		Code:           http.StatusOK,
	}
}

func (w *ResponseWriterWrapper) WriteHeader(code int) {
	w.Code = code

	w.ResponseWriter.WriteHeader(code)
}

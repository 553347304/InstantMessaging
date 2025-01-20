package middleware

import "net/http"

type Writer struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (w *Writer) Write(data []byte) (int, error) {
	w.Body = data
	return w.ResponseWriter.Write(data)
}
func (w *Writer) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
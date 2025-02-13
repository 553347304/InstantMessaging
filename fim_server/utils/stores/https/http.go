package https

import (
	"bytes"
)

func Get(url string, header map[string]string, form map[string]any) httpResponse {
	if form != nil {
		url += "?" + Form(form)
	}
	r, body, err := newRequest(setHttp{Method: "GET", URL: url, Body: nil, Header: header})
	return httpResponse{Response: r, Body: body, Error: err}
}
func Post(url string, header map[string]string, data []byte) httpResponse {
	r, body, err := newRequest(setHttp{Method: "POST", URL: url, Body: bytes.NewBuffer(data), Header: header})
	return httpResponse{Response: r, Body: body, Error: err}
}
func Put(url string, header map[string]string, data []byte) httpResponse {
	r, body, err := newRequest(setHttp{Method: "PUT", URL: url, Body: bytes.NewBuffer(data), Header: header})
	return httpResponse{Response: r, Body: body, Error: err}
}

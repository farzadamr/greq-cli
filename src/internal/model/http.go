package model

import "net/http"

type HttpRequest struct {
	Method string
	URL    string
	Header map[string]string
	Body   []byte
}

type HTTPResponse struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

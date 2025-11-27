package model

import (
	"net/http"
	"time"
)

type HTTPResponse struct {
	Method     string
	Path       string
	StatusCode int
	Headers    http.Header
	Body       []byte
	JSON       any
}

type HttpRequest struct {
	BaseURL string
	Method  string
	Path    string
	Body    any
	Headers map[string]string
}

type RequestOptions struct {
	Timeout       time.Duration
	Retries       int
	Backoff       time.Duration
	SkipTLSVerify bool
}

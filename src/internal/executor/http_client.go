package executor

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/farzadamr/greq-cli/internal/model"
)

func SendHTTPRequest(req model.HttpRequest) (*model.HTTPResponse, error) {
	var bodyReader io.Reader
	if req.Body != nil {
		bodyReader = bytes.NewBuffer(req.Body)
	}

	httpReq, err := http.NewRequest(req.Method, req.URL, bodyReader)
	if err != nil {
		return nil, err
	}

	for k, v := range req.Header {
		httpReq.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// Send Request
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &model.HTTPResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       respBody,
	}, nil
}

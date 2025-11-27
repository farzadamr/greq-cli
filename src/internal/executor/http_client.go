package executor

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/farzadamr/greq-cli/internal/model"
)

func DoRequest(req *model.HttpRequest) (*model.HTTPResponse, error) {
	var bodyReader io.Reader
	if req.Body != nil {
		switch v := req.Body.(type) {
		case io.Reader:
			bodyReader = v
		case []byte:
			bodyReader = bytes.NewReader(v)
		case string:
			strings.NewReader(v)
		default:
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}

			bodyReader = bytes.NewReader(b)
			if req.Headers == nil {
				req.Headers = make(map[string]string)
			}
			if _, exists := req.Headers["Content_Type"]; !exists {
				req.Headers["Content_Type"] = "application/json"
			}

		}

	}

	url := req.BaseURL + "/" + req.Path
	request, err := http.NewRequest(req.Method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &model.HTTPResponse{
		Path:       req.Path,
		Method:     req.Method,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       respBody,
	}, nil
}

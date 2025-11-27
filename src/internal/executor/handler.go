package executor

import (
	"time"

	"github.com/farzadamr/greq-cli/internal/model"
)

func HandleSuite(sf *model.SuiteFile, env string) (*[]model.SuiteResult, error) {
	var testResults []model.TestResponse
	var result []model.SuiteResult
	for _, suite := range sf.Suites {
		for _, t := range suite.Tests {
			httpRequest := &model.HttpRequest{
				Method:  t.Method,
				BaseURL: sf.Env[env],
				Path:    t.Path,
				Body:    t.Body,
				Headers: sf.Global.Headers,
			}

			startTime := time.Now()

			httpResponse, err := DoRequest(httpRequest)
			if err != nil {
				return nil, err
			}

			duration := time.Since(startTime)

			testResults = append(testResults, model.TestResponse{
				HTTPResponse: *httpResponse,
				Assertion:    t.Assert,
				Duration:     duration,
			})
		}
		result = append(result, model.SuiteResult{
			Tag:            suite.Tag,
			TestsResponses: testResults,
		})
	}
	return &result, nil
}
